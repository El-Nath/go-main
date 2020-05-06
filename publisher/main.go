package main

import (
	"fmt"
	"time"

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

	requestChanSend := make(chan *Request)
	nE.BindSendChan("request_subject", requestChanSend)

	i := 0
	for {

		req := Request{Id: i}
		fmt.Printf("Sending request %d\n", req.Id)
		requestChanSend <- &req
		time.Sleep(time.Second * 1)
		i = i + 1
	}

}
