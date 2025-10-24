package main

import (
	"log"
	"net"

	"github.com/NavroO/adhub/internal/auth"
	"github.com/NavroO/adhub/proto/authpb"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, auth.New())

	log.Println("✅ Auth gRPC server running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
