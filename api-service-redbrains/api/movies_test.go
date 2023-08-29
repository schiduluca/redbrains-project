package api

import (
	"context"
	"encoding/json"
	"github.com/schiduluca/api-service-redbrains/models"
	"github.com/schiduluca/api-service-redbrains/rpc/pb"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Mock MovieServiceClient for testing purposes
type MockMovieServiceClient struct{}

func (m MockMovieServiceClient) HashMovie(ctx context.Context, in *pb.MovieRequest, opts ...grpc.CallOption) (*pb.MovieResponse, error) {
	if in.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "empty movie provided")
	}
	if in.Name == "Terrible Movie" {
		return nil, status.Error(codes.InvalidArgument, "terrible movies are not supported")
	}
	return &pb.MovieResponse{HashedName: "Hashed_" + in.Name}, nil
}

func TestHashMovieName(t *testing.T) {
	assert := assert.New(t)

	// Create a mock gRPC client
	mockClient := MockMovieServiceClient{}
	movieService := NewMovieService(mockClient)

	// Prepare a test request with a valid movie
	reqBody := `{"name": "Good Movie"}`
	req := httptest.NewRequest("POST", "/hash-movie-name", strings.NewReader(reqBody))
	w := httptest.NewRecorder()

	// Call the API handler
	movieService.HashMovieName(w, req)

	// Check the response
	assert.Equal(http.StatusOK, w.Code)

	var resp models.Movie
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(err)
	assert.Equal("Hashed_Good Movie", resp.Name)

	// Prepare a test request with an invalid movie
	reqBody = `{"name": ""}`
	req = httptest.NewRequest("POST", "/hash-movie-name", strings.NewReader(reqBody))
	w = httptest.NewRecorder()

	// Call the API handler
	movieService.HashMovieName(w, req)

	// Check the response
	assert.Equal(http.StatusBadRequest, w.Code)
	assert.Equal("rpc error: code = InvalidArgument desc = empty movie provided", w.Body.String())

	// Prepare a test request with a terrible movie
	reqBody = `{"name": "Terrible Movie"}`
	req = httptest.NewRequest("POST", "/hash-movie-name", strings.NewReader(reqBody))
	w = httptest.NewRecorder()

	// Call the API handler
	movieService.HashMovieName(w, req)

	// Check the response
	assert.Equal(http.StatusBadRequest, w.Code)
	assert.Equal("rpc error: code = InvalidArgument desc = terrible movies are not supported", w.Body.String())
}
