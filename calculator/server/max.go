package main

import (
	"github.com/PetrKasmin/go-grpc/calculator/proto"
	"io"
	"log"
)

func (s *Server) Max(stream proto.CalculatorService_MaxServer) error {
	var maximum int32 = 0
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatal(err)
		}

		if number := req.Number; number > maximum {
			maximum = number
			err = stream.Send(&proto.MaxResponse{
				Result: number,
			})

			if err != nil {
				log.Fatal(err)
			}
		}

	}
}
