package main

import (
	"context"
	"fmt"
	"github.com/PetrKasmin/go-grpc/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math"
)

func (s *Server) Sqrt(ctx context.Context, in *proto.SqrtRequest) (*proto.SqrtResponse, error) {

	number := in.Number

	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Error: negative number"),
		)
	}

	return &proto.SqrtResponse{
		Result: math.Sqrt(float64(number)),
	}, nil
}
