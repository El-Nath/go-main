package main

import (
	"log"
	"net/http"
	"time"
)
import "./infra/esconfig"
import "./infra/network"

func main() {
	elas, err := esconfig.ConnectElastic("http://localhost:9200", false)
	if err != nil {
		log.Panic(err)
	}

	env := &esconfig.Env{ES: elas}

	r := network.InitRoutes(env)

	sE := &http.Server{
		Addr:           ":1338",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	_ = sE.ListenAndServe()
}
