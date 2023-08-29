package main

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/schiduluca/api-service-redbrains/config"
	"github.com/schiduluca/api-service-redbrains/router"
	"github.com/schiduluca/api-service-redbrains/rpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

func main() {
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("could not parse env variables, error=%v", err)
	}

	cc, err := grpc.Dial(cfg.CryptoServicePATH, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := pb.NewMovieServiceClient(cc)
	r := router.NewRouter(client)

	err = http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), r.GetRouter())
	if err != nil {
		log.Fatalf("starting the server failed with error: %v", err)
	}
}
