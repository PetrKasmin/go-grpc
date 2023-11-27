package main

import (
	"context"
	"github.com/PetrKasmin/go-grpc/greet/proto"
	"io"
	"log"
)

func doGreetManyTimes(c proto.GreetServiceClient) {
	stream, err := c.GreetManyTimes(context.Background(), &proto.GreetRequest{
		FirstName: "Petr",
	})

	if err != nil {
		log.Fatal(err)
	}

	for {
		mgs, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		log.Println(mgs.Result)
	}
}
