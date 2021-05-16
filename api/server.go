package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Komplementariteten/lutra"
	"github.com/Komplementariteten/lutra/db"
)

type Server struct {
	api *surveyApi
	bbg *bbg
	s   *http.Server
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
	bbg := &bbg{}
	mux.Handle("/bbg/", http.StripPrefix("/bbg/", bbg))
	mux.Handle("/survey/", http.StripPrefix("/survey/", s))
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(staticFilePath))))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	handle := &Server{api: s, s: server, bbg: bbg}

	return handle, nil
}
