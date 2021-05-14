package api

import (
	"boardgamesurvey/model"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/Komplementariteten/lutra/db"
	"github.com/Komplementariteten/lutra/util"
	"go.mongodb.org/mongo-driver/bson"
)

const knownSurveyId = "7jdksjf8238jfwajwf343.tsykp3q4wtwsa"
const sessionCookieName = "SUS_"
const surveyCollection = "LingenBrettspiel300EuroSurvey"

type surveyApi struct {
	db *db.Db
}

func startSession(w http.ResponseWriter) string {
	sessId := util.GetRandomString(128)
	sessCookie := &http.Cookie{
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		Name:     sessionCookieName,
		Path:     "/",
		Expires:  time.Now().Local().Add(1488 * time.Hour),
	}
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
		http.Error(w, "failed to read JSON in Request", http.StatusBadRequest)
	}

	content.Owner = sessionName
	bsonD := content.ToBsonD()
	filter := &bson.M{model.ContentItemOwnerName: sessionName}
	_, err = pool.Add(r.Context(), bsonD)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m, err := pool.Get(r.Context(), filter)
	bytes, err := bson.Marshal(m)
	_, err = w.Write(bytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (api *surveyApi) addTestData() {

}

func validateApiRequest(w http.ResponseWriter, r *http.Request, httpMethod string) (string, error) {
	if r.Method != http.MethodGet {
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

	resp, err := pool.Get(r.Context(), searchContent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	by, err := bson.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(by)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (api *surveyApi) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if strings.HasPrefix(r.URL.Path, "get") {
		api.get(w, r)
	}

}
