package main

import (
	"log"
	"net/http"

	"github.com/watercolor/go-websocket-gateway/internal/server"
	"github.com/watercolor/go-websocket-gateway/routers"
)

func main() {
	s := server.NewServer()
	r := routers.NewRouter()

	if err := http.ListenAndServe(s.Addr, r); err != nil {
		log.Fatalf("ListenAndServe: %s\n", err)
	}
}
