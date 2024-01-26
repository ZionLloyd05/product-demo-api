package main

import (
	"context"
	"log"
	"main/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)


	// create the handlers
	productHandler := handlers.NewProducts(logger)

	// create a new server mux & register the handlers
	serveMux := http.NewServeMux()
	serveMux.Handle("/", productHandler)

	// create a new server
	server := &http.Server{
		Addr: "127.0.0.1:9090",
		Handler: serveMux,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// start the server
	go func ()  {
		err := server.ListenAndServe()	
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Println("Recieved terminate, graceful shutdown", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30 *time.Second)
	server.Shutdown(timeoutContext)
}