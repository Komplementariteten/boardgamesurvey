package api

import (
	"boardgamesurvey/model"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Komplementariteten/lutra"
	"github.com/Komplementariteten/lutra/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	api *surveyApi
	bbg *bbg
	s   *http.Server
	ws  *voteWebSocket
}

func (s *Server) Start() {
	log.Fatal(s.s.ListenAndServe())
}

func (s *Server) Shutdown() {
	s.s.Close()
}

func CreateNew(port int, staticFilePath string, config *lutra.LutraConfig) (*Server, error) {

	db, err := db.NewConnection(context.Background(), config.MongoDbConnectionStr, config.MongoDbName)
	if err != nil {
		panic(err)
	}
	mux := http.NewServeMux()
	s := &surveyApi{db: db}
	ctx := context.Background()
	p, err := db.Collection(ctx, sessionCollection)
	if err != nil {
		panic(err)
	}
	m, err := p.Get(ctx, &bson.M{sessionField: sessionName})
	if err != nil && err != mongo.ErrNoDocuments {
		panic(err)
	}

	sessionMap := make(map[string]interface{})
	sessionMap[sessionField] = sessionName
	sessionMap[dataField] = make([]byte, 0)
	if m != nil {
		sessionMap[dataField] = primitive.Binary((m[dataField].(primitive.Binary))).Data

	} else {
		p.AddM(ctx, &bson.M{sessionField: sessionName, dataField: sessionMap[dataField]})
	}
	poll := &model.Poll{}
	pollBin := sessionMap[dataField].([]byte)
	if len(pollBin) > 0 {
		err = json.Unmarshal(pollBin, &poll)
		if err != nil {
			panic(err)
		}
	} else {
		poll.Votes = make([]*model.VoteCount, 0)
	}
	h := newHub()
	go h.run(p, sessionMap)
	ws := CreateVoteWs(db, h, poll)
	bbg := &bbg{}
	mux.Handle("/bbg/", http.StripPrefix("/bbg/", bbg))
	mux.Handle("/survey/", http.StripPrefix("/survey/", s))
	mux.Handle("/ws", ws)
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(staticFilePath))))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	handle := &Server{api: s, s: server, bbg: bbg}

	return handle, nil
}
