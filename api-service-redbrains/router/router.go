package router

import (
	"github.com/gorilla/mux"
	"github.com/schiduluca/api-service-redbrains/api"
	"github.com/schiduluca/api-service-redbrains/rpc/pb"
	"net/http"
)

type Router struct {
	cryptoServiceClient pb.MovieServiceClient
	apiRouter           *mux.Router
}

func NewRouter(cryptoServiceClient pb.MovieServiceClient) Router {
	router := mux.NewRouter()

	movieService := api.NewMovieService(cryptoServiceClient)
	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	apiRouter.HandleFunc("/hash-movie-name", movieService.HashMovieName).Methods(http.MethodPost)

	return Router{
		apiRouter:           apiRouter,
		cryptoServiceClient: cryptoServiceClient,
	}
}

func (r Router) GetRouter() *mux.Router {
	return r.apiRouter
}

func (r Router) GetCryptoServiceClient() pb.MovieServiceClient {
	return r.cryptoServiceClient
}
