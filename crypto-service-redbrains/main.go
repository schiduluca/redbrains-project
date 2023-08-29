package main

import (
	"crypto-service-redbrains/config"
	"crypto-service-redbrains/rpc/pb"
	"crypto-service-redbrains/server"
	"fmt"
	"github.com/caarlos0/env"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("could not parse env variables, error=%v", err)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		log.Fatalf("could not listen to port, error:%v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMovieServiceServer(s, &server.Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
