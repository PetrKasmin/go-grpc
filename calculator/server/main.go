package main

import (
	"github.com/PetrKasmin/go-grpc/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Server struct {
	proto.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Listen 0.0.0.0:50051")

	srv := grpc.NewServer()
	proto.RegisterCalculatorServiceServer(srv, &Server{})
	reflection.Register(srv)

	if err = srv.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
