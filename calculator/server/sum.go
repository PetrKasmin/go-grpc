package main

import (
	"context"
	"github.com/PetrKasmin/go-grpc/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *proto.SumRequest) (*proto.SumResponse, error) {
	return &proto.SumResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}
