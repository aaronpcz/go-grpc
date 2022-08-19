package main

import (
	"context"
	"io"
	"log"

	pb "github.com/aaronpcz/go-grpc/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("doPrimes was invoked")

	req := &pb.PrimesRequest{
		Primes: 120,
	}

	stream, err := c.Primes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

		// inf loop
		for {
			msg, err := stream.Recv()
	
			// stream is over
			if err == io.EOF {
				break
			}
	
			if err != nil {
				log.Fatalf("Error while reading stream: %v\n", err)
			}
	
			log.Printf("Primes: %d\n", msg.Rsult)
		}
}