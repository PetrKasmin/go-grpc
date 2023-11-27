package main

import (
	"fmt"
	"github.com/PetrKasmin/go-grpc/greet/proto"
)

func (s *Server) GreetManyTimes(in *proto.GreetRequest, stream proto.GreetService_GreetManyTimesServer) error {
	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i)

		stream.Send(&proto.GreetResponse{
			Result: res,
		})
	}

	return nil
}
