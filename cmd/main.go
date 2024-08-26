package main

import (
	"github.com/watercolor/go-websocket-gateway/internal/server"
	"github.com/watercolor/go-websocket-gateway/routers"
)

func main() {
	s := server.NewServer()
	r := routers.NewRouter()

	s.Run(r)
}
