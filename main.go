package main

import (
	"appointy-task/api"
	"appointy-task/db"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	server := &http.Server{Addr: ":8080", Handler: api.RouteHandler{}}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	interruptSignal := make(chan os.Signal, 1)
	signal.Notify(interruptSignal, os.Interrupt)
	<-interruptSignal

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	db.Close(ctx)
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln(err)
	}
}
