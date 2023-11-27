package main

import (
	"github.com/PetrKasmin/go-grpc/calculator/proto"
)

func (s *Server) Primes(in *proto.PrimeRequest, stream proto.CalculatorService_PrimesServer) error {

	var k int64 = 2
	n := in.Number

	for n > 1 {
		if n%k == 0 {
			stream.Send(&proto.PrimeResponse{
				Result: k,
			})

			n = n / k
		} else {
			k = k + 1
		}
	}

	return nil
}
