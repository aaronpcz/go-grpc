package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/aaronpcz/go-grpc/greet/proto"
)

var addr string = "localhost:8080"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()
	// .. do somethings, at end, defer will close
	
	c := pb.NewGreetServiceClient(conn)
	// doGreet(c)
	doGreetManyTimes(c)
}