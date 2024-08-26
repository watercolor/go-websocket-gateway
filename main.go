package main

import (
	"log"
	"net/http"

	"github.com/watercolor/go-websocket-gateway/internal/server"
	"github.com/watercolor/go-websocket-gateway/pkg/logger"
	"github.com/watercolor/go-websocket-gateway/routers"
)

func main() {
	s := server.NewServer()
	r := routers.NewRouter()
	logger.Info("Server start..")
	if err := http.ListenAndServe(s.Addr, r); err != nil {
		log.Fatalf("ListenAndServe: %s\n", err)
	}
}
