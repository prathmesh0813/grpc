package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/prathmesh0813/grpc/calculator/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr,grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!= nil {
		log.Fatalf("Failed to connect: %v\n",err)
	}
	client := pb.NewCalculatorServiceClient(conn)
	doSum(client)
	defer conn.Close()
}