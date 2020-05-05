package main

import (
	"log"
	"net/http"
	"time"

	"./infra/network"
	"./infra/redconfig"
)

func main() {

	cR, err := redconfig.ConnectRedis("localhost:6379", "", 0)
	if err != nil {
		log.Panic(err)
	}

	env := &redconfig.Env{RC: cR}

	r := network.InitRoutes(env)

	rD := &http.Server{
		Addr:           ":1339",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	_ = rD.ListenAndServe()

}
