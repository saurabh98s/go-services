package main

import (
	"context"
	"log"
	"micro-services/handlers"
	"os/signal"
	"time"

	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product-api ", log.LstdFlags)
	ph:=handlers.NewProducts(l)
	sm := http.NewServeMux()
	// sm.Handle("/", hh)
	// sm.Handle("/goodbye", gh)
	sm.Handle("/",ph)

	s := http.Server{
		Addr:         ":9090", //binding the address
		Handler:      sm, //setting the default address
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// start the server
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	//traps sigterm or interupt and gracefully shutdown the server	
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Recieved terminate signal, graceful shutdown\nsignal type:", sig)

	// Graceful shutdown
	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(timeoutContext)

}
