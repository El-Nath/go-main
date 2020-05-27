package main

import (
	"fmt"

	nats "github.com/nats-io/nats.go"
)

type Request struct {
	Id int
}

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}

	nE, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		panic(err)
	}
	defer nE.Close()

	requestChanRecv := make(chan *Request)
	nE.BindRecvChan("request_subject", requestChanRecv)

	for {
		req := <-requestChanRecv
		fmt.Printf("Received request: %d\n", req.Id)
	}
}
