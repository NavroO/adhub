package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/NavroO/adhub/internal/auth"
	"github.com/NavroO/adhub/internal/shared"
	"github.com/NavroO/adhub/proto/authpb"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {
	shared.SetupLogger()
	log.Info().Msg("📦 Logger initialized")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}

	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, auth.New())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Info().Msg("✅ Auth gRPC server running on :50051")
		if err := s.Serve(lis); err != nil {
			log.Fatal().Err(err).Msg("failed to serve")
		}
	}()

	<-quit
	log.Info().Msg("🧹 Shutting down server...")

	s.GracefulStop()

	log.Info().Msg("✅ Server exited gracefully")
}
