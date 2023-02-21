package main

import (
	"context"
	"errors"
	"github.com/AthulMuralidhar/my-site-2/api"
	"github.com/AthulMuralidhar/my-site-2/server"
	"log"
	"net"
	"net/http"
)

// https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", api.GetIndex)

	ctx, cancel := context.WithCancel(context.Background())
	defaultServer := &http.Server{
		Addr:    ":8000",
		Handler: mux,
		BaseContext: func(listener net.Listener) context.Context {
			ctx = context.WithValue(ctx, server.DefaultServerName, listener.Addr().String())
			return ctx
		},
	}
	go func() {
		err := defaultServer.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			log.Println("Default server closed")
		} else if err != nil {
			log.Printf("Default server error: ", err)
		}
		cancel()
	}()
	log.Printf("Default server started at: %s", defaultServer.Addr)
	<-ctx.Done()
}
