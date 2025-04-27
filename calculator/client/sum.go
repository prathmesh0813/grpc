package main

import (
	"context"
	"log"

	pb "github.com/prathmesh0813/grpc/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) (){
	log.Println("doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		Num1: 1,Num2: 2,
	})

	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("Sum: %d\n",res.Result)
}