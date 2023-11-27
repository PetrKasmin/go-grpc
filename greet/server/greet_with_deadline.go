package main

import (
	"context"
	"errors"
	"github.com/PetrKasmin/go-grpc/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *proto.GreetRequest) (*proto.GreetResponse, error) {
	for i := 0; i < 3; i++ {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			msg := "Client canceled request"
			log.Println(msg)
			return nil, status.Error(codes.Canceled, msg)
		}

		time.Sleep(1 * time.Second)
	}

	return &proto.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
