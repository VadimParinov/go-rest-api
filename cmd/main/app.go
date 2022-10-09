package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"rest-api/cmd/internal/user"
	"time"
)

func main() {
	log.Println("Starting server...")
	router := httprouter.New()
	handler := user.NewHandler()
	log.Println("Registering handlers...")
	handler.Register(router)
	start(router)

}

func start(router *httprouter.Router) {
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.Serve(listener))
}
