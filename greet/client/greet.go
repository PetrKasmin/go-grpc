package main

import (
	"context"
	pb "github.com/PetrKasmin/go-grpc/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked!")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Petr",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %v\n", res.Result)
}
