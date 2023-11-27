package main

import (
	"context"
	"github.com/PetrKasmin/go-grpc/blog/proto"
	"log"
)

func readBlog(c proto.BlogServiceClient, id string) *proto.Blog {
	log.Println("---readBlog was invoked---")

	req := &proto.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Fatalf("Error happened while reading: %v\n", err)
	}

	log.Printf("Blog was read: %v\n", res)
	return res
}
