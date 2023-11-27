package main

import (
	"context"
	"github.com/PetrKasmin/go-grpc/greet/proto"
	"log"
	"time"
)

func doLongGreet(c proto.GreetServiceClient) {
	reqs := []*proto.GreetRequest{
		{FirstName: "Petr"},
		{FirstName: "Anton"},
		{FirstName: "Ignat"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range reqs {
		log.Printf("Sending %s\n", r.FirstName)
		stream.Send(r)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.Result)

}
