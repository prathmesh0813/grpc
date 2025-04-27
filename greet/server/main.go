package main

import (
	"log"
	"net"

	pb "github.com/prathmesh0813/grpc/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct{
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp",addr)
	if err != nil {
		log.Fatalf("Failed to listen on : %v\n", err)
	}

	log.Printf("listening on %s\n",addr)

	s:= grpc.NewServer()

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}