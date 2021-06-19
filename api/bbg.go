package api

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type bbg struct {
}

type BoardGame struct {
	ObjectId  int    `xml:"objectid,attr"`
	Name      string `xml:"name"`
	Published string `xml:"yearpublished"`
}

type Results struct {
	XMLName xml.Name    `xml:"boardgames"`
	Games   []BoardGame `xml:"boardgame"`
}

type Players struct {
	Value int `xml:"value,attr"`
}

type Name struct {
	Value string `xml:"value,attr"`
	Type  string `xml:"type,attr"`
	Order string `xml:"sortindex,attr"`
}

type BbgLink struct {
	Type  string `xml:"type,attr"`
	Id    int    `xml:"id,attr"`
	Value string `xml:"value,attr"`
}

type Thing struct {
	Type        string    `xml:"type,attr"`
	ObjectId    int       `xml:"id,attr"`
	MinPlayers  Players   `xml:"minplayers"`
	MaxPlayers  Players   `xml:"maxplayers"`
	Names       []Name    `xml:"name"`
	Thumb       string    `xml:"thumbnail"`
	Image       string    `xml:"image"`
	Description string    `xml:"description"`
	Links       []BbgLink `xml:"link"`
}

type GameThing struct {
	Type        string
	ObjectId    int
	MinPlayers  int
	MaxPlayers  int
	Image       string
	Thumb       string
	Description string
	Tags        []BbgLink
	Names       []string
}

type Things struct {
	XMLName xml.Name `xml:"items"`
	Things  []Thing  `xml:"item"`
}

const apiTextSearchParam = "name"
const bbgApiSearchUrl = "https://www.boardgamegeek.com/xmlapi/search?search"
const bbgApiThingUrl = "https://www.boardgamegeek.com/xmlapi2/thing?id"

func (t *Thing) ToGameThing() *GameThing {
	res := &GameThing{}
	res.Description = t.Description
	res.Type = t.Type
	res.Thumb = t.Thumb
	res.Image = t.Thumb
	res.MinPlayers = t.MinPlayers.Value
	res.MaxPlayers = t.MaxPlayers.Value
	res.ObjectId = t.ObjectId
	res.Tags = t.Links
	numOfNames := len(t.Names)
	res.Names = make([]string, numOfNames)
	for i := 0; i < numOfNames; i++ {
		res.Names[i] = t.Names[i].Value
	}
	return res
}

func (bbg *bbg) searchBBG(searchParams map[string][]string) (*Results, error) {
	if name, ok := searchParams[apiTextSearchParam]; ok {
		reqUrl := fmt.Sprintf("%s=%s", bbgApiSearchUrl, url.QueryEscape(name[0]))
		resp, err := http.Get(reqUrl)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			result := &Results{}
			err = xml.Unmarshal(bodyBytes, &result)
			if err != nil {
				return nil, err
			}
			return result, nil
		}

	}
	return nil, errors.New("searchParam not found")
}

func (bbg *bbg) getThing(id int64) (*Thing, error) {
	reqUrl := fmt.Sprintf("%s=%d", bbgApiThingUrl, id)
	resp, err := http.Get(reqUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &Things{}
		err = xml.Unmarshal(bodyBytes, &result)
		if err != nil {
			return nil, err
		}
		return &result.Things[0], nil
	}
	return nil, errors.New("request not OK")
}

func (bbg *bbg) handleThing(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 2 {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	intId, err := strconv.ParseInt(parts[1], 0, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	thing, err := bbg.getThing(intId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bytes, err := json.Marshal(thing.ToGameThing())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (bbg *bbg) handleSearch(w http.ResponseWriter, r *http.Request) {
	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		panic(err)
	}
	results, err := bbg.searchBBG(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	games := results.Games

	var bytes []byte
	if games != nil {
		bytes, err = json.Marshal(games)
	} else {
		d := make([]int, 0)
		bytes, err = json.Marshal(d)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (bbg *bbg) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Not implemented", http.StatusMethodNotAllowed)
	}

	if strings.HasPrefix(r.URL.Path, "search") {
		bbg.handleSearch(w, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "game") {
		bbg.handleThing(w, r)
		return
	}
}
