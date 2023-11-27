package main

import (
	"context"
	"github.com/PetrKasmin/go-grpc/calculator/proto"
	"io"
	"log"
)

func doMax(c proto.CalculatorServiceClient) {
	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	nums := []int32{1, 5, 3, 6, 2, 20}

	ch := make(chan struct{})

	go func() {
		for _, n := range nums {
			stream.Send(&proto.MaxRequest{
				Number: n,
			})
		}
	}()

	go func() {
		defer close(ch)
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatal(err)
			}

			log.Println(res.Result)
		}
	}()

	<-ch
}
