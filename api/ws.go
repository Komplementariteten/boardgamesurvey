package api

import (
	"boardgamesurvey/model"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Komplementariteten/lutra/db"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 2048
)
const sessionCollection = "LingenBrettspiel300EuroSession"
const sessionName = "bgsurvey"
const sessionField = "session"
const dataField = "data"

type voteWebSocket struct {
	db   *db.Db
	poll *model.Poll
	hub  *hub
	ctrl chan bool
}

type hub struct {
	// Registered clients.
	clients map[*client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *client

	// Unregister requests from clients.
	unregister chan *client
}

type client struct {
	hub *hub
	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

type state struct {
	data []byte
}

var (
	newline  = []byte{'\n'}
	space    = []byte{' '}
	upgrader = websocket.Upgrader{
		ReadBufferSize:  4086,
		WriteBufferSize: 4086,
	}
)

func newHub() *hub {
	return &hub{
		broadcast:  make(chan []byte),
		register:   make(chan *client),
		unregister: make(chan *client),
		clients:    make(map[*client]bool),
	}
}

func (h *hub) run(p *db.Pool, poll map[string]interface{}) {
	ctx := context.Background()
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			if len(poll[dataField].([]byte)) > 0 {
				client.send <- poll[dataField].([]byte)
			}
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			poll[dataField] = message
			_, err := p.Update(ctx, &bson.M{sessionField: sessionName}, poll)
			if err != nil {
				fmt.Printf("Database update error: %v\n", err)
			}
			fmt.Printf("got message %v\n", message)
			for client := range h.clients {
				fmt.Println("notifieing client")
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

func CreateVoteWs(db *db.Db, h *hub, poll *model.Poll) *voteWebSocket {
	ws := &voteWebSocket{
		db:   db,
		poll: poll,
		ctrl: make(chan bool),
		hub:  h,
	}
	return ws
}

func (c *client) wsReader(ws *voteWebSocket) {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	vote := model.Vote{}
	for {
		msgType, p, err := c.conn.ReadMessage()
		fmt.Printf("%v : %v", msgType, string(p))
		json.Unmarshal(p, &vote)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v\n", err)
			}
			break
		}
		var isNew bool = true
		for _, v := range ws.poll.Votes {
			var hasVoted bool
			hasVoted = false
			if v.Title == vote.Title {
				isNew = false
				for _, o := range v.Voters {
					if o == vote.Owner {
						hasVoted = true
					}
				}
			}
			if !hasVoted {
				v.Votes++
				v.Voters = append(v.Voters, vote.Owner)
			}
		}
		if isNew {
			newVote := &model.VoteCount{
				Title:    vote.Title,
				Votes:    1,
				ObjectId: vote.ObjectId,
				Voters:   []string{vote.Owner},
			}
			ws.poll.Votes = append(ws.poll.Votes, newVote)
		}
		bytes, err := json.Marshal(ws.poll)
		if err != nil {
			log.Printf("error: %v\n", err)
		}
		c.hub.broadcast <- bytes
	}
}

func (c *client) wsWriter(ws *voteWebSocket) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		case <-ws.ctrl:
			log.Println("closing")
			c.conn.Close()
		}
	}
}

func (ws *voteWebSocket) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Origin") != "http://"+r.Host {
		http.Error(w, "Origin not allowed", 403)
		return
	}
	fmt.Println("Try to connect new WS client")
	conn, err := upgrader.Upgrade(w, r, w.Header())
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	client := &client{hub: ws.hub, conn: conn, send: make(chan []byte, 512)}
	ws.hub.register <- client
	go client.wsReader(ws)
	go client.wsWriter(ws)
	fmt.Println("client connected")
}
