package main

import (
	"log"

	pb "github.com/aaronpcz/go-grpc/calculator/proto"
)

func (s *Server) Primes(in *pb.PrimesRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes function was invoked with %v\n", in)

	k := int32(2)
	N := in.Primes

	for {

		if N <= 1 {
			break
		}

		if N % k == 0 {
			stream.Send(&pb.PrimesResponse{
				Rsult: k,
			})
			N = N / k
		} else {
			k = k + 1
		}
			
	}

	return nil
}