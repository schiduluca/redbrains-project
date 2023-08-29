package api

import (
	"encoding/json"
	"errors"
	"github.com/schiduluca/api-service-redbrains/models"
	"github.com/schiduluca/api-service-redbrains/rpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
)

type MovieService struct {
	cryptoServiceClient pb.MovieServiceClient
}

func NewMovieService(client pb.MovieServiceClient) MovieService {
	return MovieService{
		cryptoServiceClient: client,
	}
}

func (m MovieService) HashMovieName(writer http.ResponseWriter, reader *http.Request) {
	ctx := reader.Context()
	var movie models.Movie
	err := json.NewDecoder(reader.Body).Decode(&movie)
	if err != nil {
		log.Println("could not encode the response")
		writeErrorResponse(writer, errors.New("invalid request provided"))
		return
	}

	hashMovie, err := m.cryptoServiceClient.HashMovie(ctx, &pb.MovieRequest{Name: movie.Name})
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.InvalidArgument {
			writeErrorResponse(writer, st.Err())
			return
		}
		log.Printf("internal error happened, err=%v", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := models.Movie{
		Name: hashMovie.HashedName,
	}
	err = json.NewEncoder(writer).Encode(&resp)
	if err != nil {
		log.Print("could not encode the response")
	}
}

func writeErrorResponse(writer http.ResponseWriter, err error) {
	writer.WriteHeader(http.StatusBadRequest)
	writer.Write([]byte(err.Error()))
}
