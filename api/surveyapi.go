package api

import (
	"boardgamesurvey/model"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Komplementariteten/lutra/db"
	"github.com/Komplementariteten/lutra/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const knownSurveyId = "7jdksjf8238jfwajwf343"
const sessionCookieName = "SUS_"
const surveyCollection = "LingenBrettspiel300EuroSurvey"
const sessionCollection = "LingenBrettspiel300EuroSession"

type surveyApi struct {
	db *db.Db
}

func startSession(w http.ResponseWriter) string {
	sessId := util.GetRandomString(128)
	sessCookie := &http.Cookie{
		HttpOnly: false,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		Name:     sessionCookieName,
		Path:     "/",
		Expires:  time.Now().Local().Add(1488 * time.Hour),
	}
	sessCookie.Value = sessId
	http.SetCookie(w, sessCookie)
	return sessId
}

func (api *surveyApi) set(w http.ResponseWriter, r *http.Request) {
	sessionName, err := validateApiRequest(w, r, http.MethodPost)
	if err != nil {
		return
	}
	pool, err := api.db.Collection(r.Context(), surveyCollection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, content := model.CheckJsonInHttp(r)
	if resp.Error {
		http.Error(w, resp.Message, http.StatusBadRequest)
		return
	}

	content.Owner = sessionName
	bsonD := content.ToBsonD()
	filter := &bson.M{model.ContentItemOwnerName: sessionName}
	_, err = pool.Get(r.Context(), filter)
	if err == mongo.ErrNoDocuments {
		_, err = pool.Add(r.Context(), bsonD)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = pool.Update(r.Context(), filter, bsonD.Map())

	if err != nil && err.Error() != "No Changes where made" {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Write(w)
}

func validateApiRequest(w http.ResponseWriter, r *http.Request, httpMethod string) (string, error) {
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Not supported", http.StatusNotImplemented)
		return "", errors.New("Not supported")
	}
	surveyIdIndex := strings.LastIndex(r.URL.Path, "/")
	surveyId := r.URL.Path[surveyIdIndex+1:]
	if surveyId != knownSurveyId {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return "", errors.New("forbidden")
	}

	var sessionName string
	c, err := r.Cookie(sessionCookieName)
	if err == nil {
		sessionName = c.Value
	}
	if err == http.ErrNoCookie {
		sessionName = startSession(w)
		return sessionName, nil
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "", err
	}
	return sessionName, nil
}

func (api *surveyApi) get(w http.ResponseWriter, r *http.Request) {
	sessionName, err := validateApiRequest(w, r, http.MethodGet)
	if err != nil {
		return
	}

	pool, err := api.db.Collection(r.Context(), surveyCollection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	searchContent := &bson.M{model.ContentItemOwnerName: sessionName}

	m, err := pool.Get(r.Context(), searchContent)
	fmt.Printf("%v\n", m)
	if err == mongo.ErrNoDocuments {
		w.Write([]byte("{}"))
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := model.JsonAsResponse(m)
	resp.Write(w)
}

func (api *surveyApi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		api.get(w, r)
	}
	if r.Method == http.MethodPost {
		api.set(w, r)
	}
}
