package main

import (
	"github.com/PetrKasmin/go-grpc/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := proto.NewCalculatorServiceClient(conn)
	//doSum(c)
	//doPrimes(c)
	//doAvg(c)
	//doMax(c)
	doSqrt(c, 10)
}
