package main

import (
	"context"
	"github.com/PetrKasmin/go-grpc/calculator/proto"
	"io"
	"log"
)

func doPrimes(c proto.CalculatorServiceClient) {
	stream, err := c.Primes(context.Background(), &proto.PrimeRequest{
		Number: 120,
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
