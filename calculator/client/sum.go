package main

import (
	"context"
	"log"
	"math/rand"

	pb "github.com/aaronpcz/go-grpc/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		A: int32(rand.Intn(10)),
		B: int32(rand.Intn(10)),
	})

	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("Sum: %d\n", res.Result)
}