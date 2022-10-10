package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"rest-api/cmd/internal/config"
	"rest-api/cmd/internal/user"
	"time"
)

func main() {
	log.Println("Starting server...")
	router := httprouter.New()
	handler := user.NewHandler()
	cfg := config.GetConfig()
	log.Println("Registering handlers...")
	handler.Register(router)
	start(router, cfg)

}

func start(router *httprouter.Router, cfg *config.Config) {
	var listener net.Listener
	var ListenerErr error

	if cfg.Listen.Type == "sock" {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}
		socketPath := path.Join(appDir, "app.sock")
		listener, ListenerErr = net.Listen("unix", socketPath)
		if ListenerErr != nil {
			log.Fatal(ListenerErr)
		}
		log.Println("Server started on socket: ", socketPath)
	} else {
		listener, ListenerErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Port))
		if ListenerErr != nil {
			log.Fatal(ListenerErr)
		}
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.Serve(listener))
}
