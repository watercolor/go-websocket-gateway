package routers

import (
	"github.com/gorilla/mux"
	"github.com/watercolor/go-websocket-gateway/internal/handlers"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/gateway/ws", handlers.WebSocketHandler).Methods("GET")

	return router
}
