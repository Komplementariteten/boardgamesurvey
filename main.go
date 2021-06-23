package main

import (
	"boardgamesurvey/api"
	"fmt"

	"github.com/Komplementariteten/lutra"
)

func main() {

	cfg := &lutra.LutraConfig{
		MongoDbConnectionStr: "mongodb://localhost/defaultDatabase?retryWrites=true&w=majority",
		//MongoDbConnectionStr: "mongodb://mongo-host/defaultDatabase?retryWrites=true&w=majority",
		MongoDbName: "lingen-boardgame-survey",
	}

	s, err := api.CreateNew(6100, "./static/", cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting API Server")
	s.Start()

}
