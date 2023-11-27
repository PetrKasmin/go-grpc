package main

import (
	"context"
	"github.com/PetrKasmin/go-grpc/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func doGreetWithDeadline(c proto.GreetServiceClient, timeout time.Duration) {

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := c.GreetWithDeadline(ctx, &proto.GreetRequest{
		FirstName: "Petr",
	})

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println(e.Code(), e.Message())
			} else {
				log.Fatal(err)
			}
			return
		} else {
			log.Fatal(err)
		}
	}

	log.Println(res.Result)
}
