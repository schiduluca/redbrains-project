package server

import (
	"context"
	"crypto-service-redbrains/rpc/pb"
	"crypto-service-redbrains/utils"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

type Server struct {
	pb.MovieServiceServer
}

func (s *Server) HashMovie(ctx context.Context, req *pb.MovieRequest) (*pb.MovieResponse, error) {
	if err := ValidateMovie(req.Name); err != nil {
		return &pb.MovieResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	return &pb.MovieResponse{
		HashedName: utils.NewSHA256String(req.Name),
	}, nil
}

func ValidateMovie(movie string) error {
	if movie == "" {
		return errors.New("empty movie provided")
	}

	if strings.Contains(strings.ToLower(movie), "fast and furious") {
		return errors.New("terrible movies are not supported")
	}

	return nil
}
