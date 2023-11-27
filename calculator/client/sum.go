package main

import (
	"context"
	"github.com/PetrKasmin/go-grpc/calculator/proto"
	"log"
)

func doSum(c proto.CalculatorServiceClient) {
	res, err := c.Sum(context.Background(), &proto.SumRequest{
		FirstNumber:  20,
		SecondNumber: 50,
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.Result)
}
