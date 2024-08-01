package main

import (
	"ddd-implementation/pkg/offersearch/application/handler"
	"ddd-implementation/pkg/offersearch/domain/service/getoffer"
	"ddd-implementation/pkg/offersearch/domain/service/saveoffer"
	"ddd-implementation/pkg/offersearch/infrastructure/config"
	"ddd-implementation/pkg/offersearch/infrastructure/mysql/ods"
	"ddd-implementation/pkg/offersearch/infrastructure/web/direktanbindung"
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
	getOfferService := getoffer.NewService(daClient)
	saveOfferService := saveoffer.NewService(odsOfferRepo)

	// create handlers
	offerHandler := handler.NewOfferHandler(getOfferService, saveOfferService)

	mux := http.NewServeMux()

	// register all handlers
	mux.HandleFunc("/getoffers", offerHandler.Handle)

	return mux
}
