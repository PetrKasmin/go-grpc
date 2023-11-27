package main

import (
	"context"
	"github.com/PetrKasmin/go-grpc/greet/proto"
	"io"
	"log"
	"time"
)

func doGreetEveryone(c proto.GreetServiceClient) {
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	reqs := []*proto.GreetRequest{
		{FirstName: "Petr"},
		{FirstName: "Anton"},
		{FirstName: "Ignat"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Println(req.FirstName)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatal(err)
			}

			log.Println(res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
