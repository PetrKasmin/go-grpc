package main

import (
	"github.com/PetrKasmin/go-grpc/calculator/proto"
	"io"
	"log"
)

func (s *Server) Avg(stream proto.CalculatorService_AvgServer) error {
	var sum int32 = 0
	count := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&proto.AvgResponse{
				Result: float64(sum) / float64(count),
			})
		}

		if err != nil {
			log.Fatal(err)
		}

		sum += req.Number
		count++
	}

	return nil
}
