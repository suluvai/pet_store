package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"pet_store_rest_api/authentication"
	"pet_store_rest_api/controllers"
	"pet_store_rest_api/repositories"
	"pet_store_rest_api/restapi"

	"github.com/gorilla/mux"
)

func main() {

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15,
		"the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")

	var listenAddress string
	flag.StringVar(&listenAddress, "listen-addr", ":9000", "server listen address")

	var accessKey string
	flag.StringVar(&accessKey, "access-key", "PetStoreAccessKey", "api access key")

	flag.Parse()

	app := &restapi.App{
		Router: mux.NewRouter().StrictSlash(true),
	}

	petController := controllers.PetController{repositories.PetRepoData}
	auth := authentication.Secret{accessKey}

	app.SetupRouter(&petController, &auth)

	log.Println("Starting Server on ", listenAddress)

	srv := &http.Server{
		Handler:      app.Router,
		Addr:         listenAddress,
		WriteTimeout: wait,
		ReadTimeout:  wait,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
