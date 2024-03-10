package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alexispell/creep/gateway/internal/config"
	"github.com/alexispell/creep/gateway/internal/server/router"
)

type Server struct {
	*config.Config
	Router *http.ServeMux
}

func NewServer(config *config.Config) *Server {
	return &Server{
    Config: config,
	}
}

func (s *Server) RegisterRouter() {
	s.Router = router.RegisterRouter()
}

func (s *Server) StartServer() {
	fmt.Printf("Starting server on port %d\n", s.Config.Port)

	addr := fmt.Sprintf(":%d", s.Config.Port)

	err := http.ListenAndServe(addr, s.Router)
	if err != nil {
		log.Fatalf("Error starting server :>> %s", err)
	}
}
