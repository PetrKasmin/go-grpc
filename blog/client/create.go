package main

import (
	"context"
	"github.com/PetrKasmin/go-grpc/blog/proto"
	"log"
)

func createBlog(c proto.BlogServiceClient) string {
	log.Println("---createBlog was invoked---")

	blog := &proto.Blog{
		AuthorId: "Clement",
		Title:    "My First Blog",
		Content:  "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %v\n", res)
	return res.Id
}
