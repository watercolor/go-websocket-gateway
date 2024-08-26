package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/watercolor/go-websocket-gateway/config"
)

type Server struct {
	Addr string
	*http.Server
}

func NewServer() *Server {
	cfg := config.GetConfig()
	return &Server{
		Addr: cfg.App.Host + ":" + cfg.App.Port,
		Server: &http.Server{
			Addr:    cfg.App.Host + ":" + cfg.App.Port,
			Handler: nil, // 初始化为空，将在 Run 方法中设置
		},
	}
}

func (s *Server) Run(router *mux.Router) {
	s.Handler = router

	log.Printf("Starting server on %s", s.Addr)
	if err := s.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe: %s\n", err)
	}
}
