package main

import (
	"context"
	"github.com/PetrKasmin/go-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var collection *mongo.Collection

type Server struct {
	proto.BlogServiceServer
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Listen 0.0.0.0:50051")

	srv := grpc.NewServer()
	proto.RegisterBlogServiceServer(srv, &Server{})
	reflection.Register(srv)

	if err = srv.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
