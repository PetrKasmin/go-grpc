package main

import (
	"fmt"
	"github.com/PetrKasmin/go-grpc/greet/proto"
	"io"
	"log"
)

func (s *Server) GreetEveryone(stream proto.GreetService_GreetEveryoneServer) error {
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatal(err)
		}

		res := fmt.Sprintf("Hello %s!", req.FirstName)
		err = stream.Send(&proto.GreetResponse{
			Result: res,
		})
		if err != nil {
			log.Fatal(err)
		}
	}

}
