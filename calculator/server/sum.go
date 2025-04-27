package main

import (
	"context"
	"log"

	pb "github.com/prathmesh0813/grpc/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse,error){
	log.Printf("Greet function was invoked with %v\n",in)
	return &pb.SumResponse{
		Result: in.Num1+ in.Num2,
	},nil
}