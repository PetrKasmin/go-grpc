package main

import (
	"context"
	"github.com/PetrKasmin/go-grpc/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func doSqrt(c proto.CalculatorServiceClient, n int32) {
	res, err := c.Sqrt(context.Background(), &proto.SqrtRequest{
		Number: n,
	})

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Println(e.Code(), e.Message())

			if e.Code() == codes.InvalidArgument {
				log.Println(e.String(), e.Details())
			}
			return
		} else {
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	log.Println(res)
}
