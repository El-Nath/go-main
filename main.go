package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"./infra/config"
	"./infra/network"
	_ "github.com/lib/pq"
)

func main() {
	db, err := config.NewDB("postgres", "postgres://postgres:bijigaban13@172.17.0.5:5432/postgres?sslmode=disable")
	if err != nil {
		log.Panic(err)
	}

	env := &config.Env{DB: db}

	r := network.InitRoutes(env)

	s := &http.Server{
		Addr:           ":1441",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("running on port", s.Addr)
	_ = s.ListenAndServe()
}
