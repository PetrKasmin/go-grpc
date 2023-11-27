package main

import (
	"context"
	"github.com/PetrKasmin/go-grpc/blog/proto"
	"log"
)

func updateBlog(c proto.BlogServiceClient, id string) {
	log.Println("---updateBlog was invoked---")
	newBlog := &proto.Blog{
		Id:       id,
		AuthorId: "Changed Author",
		Title:    "My First Blog (edited)",
		Content:  "Content of the first blog, with some awesome additions!",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Printf("Error happened while updating: %v\n", err)
	}

	log.Println("Blog was updated")
}
