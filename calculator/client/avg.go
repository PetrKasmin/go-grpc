package main

import (
	"context"
	"github.com/PetrKasmin/go-grpc/calculator/proto"
	"log"
)

func doAvg(client proto.CalculatorServiceClient) {
	reqs := []int32{1, 2, 3, 4}

	stream, err := client.Avg(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, n := range reqs {
		stream.Send(&proto.AvgRequest{
			Number: n,
		})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.Result)
}
