package main

import (
	"context"
	"log"

	pb "github.com/aaronpcz/go-grpc/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Aaron",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}