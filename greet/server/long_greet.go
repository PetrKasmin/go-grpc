package main

import (
	"fmt"
	"github.com/PetrKasmin/go-grpc/greet/proto"
	"io"
	"log"
)

func (s *Server) LongGreet(stream proto.GreetService_LongGreetServer) error {
	res := ""
	for {

		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&proto.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatal(err)
		}

		res += fmt.Sprintf("Hello %s\n", req.FirstName)

	}

	return nil
}
