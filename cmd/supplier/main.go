package main

import (
	"awesomeProject/pkg/offersearch/application/handler"
	"awesomeProject/pkg/offersearch/domain/service/manipulateOffer"
	"awesomeProject/pkg/offersearch/infrastructure/config"
	"awesomeProject/pkg/offersearch/infrastructure/mysql/ods"
	"awesomeProject/pkg/offersearch/infrastructure/web/direktanbindung"
	"fmt"
	"net/http"
	"os"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}

func run() error {
	srv := bootstrap()

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: srv,
	}

	fmt.Println("starting server listening on port 8080")

	if err := httpServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func bootstrap() *http.ServeMux {
	// load config from .env or whatever
	conf := config.NewConfig()
	// create http clients
	daClient := direktanbindung.NewClient(conf.DaBaseUrl)

	// create repositories
	odsOfferRepo := ods.NewOfferRepository()

	// create services
	manipulateOfferService := manipulateOffer.NewService(odsOfferRepo, daClient)

	// create handlers
	offerHandler := handler.NewOfferHandler(manipulateOfferService)

	mux := http.NewServeMux()

	// register all handlers
	mux.HandleFunc("/getoffers", offerHandler.Handle)

	return mux
}