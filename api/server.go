package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/urfave/negroni"
	"github.com/vboluda/go-backend-task-management/config"
)

type Server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	s := new(Server)
	s.cfg = cfg
	return s
}

func (s *Server) Start() {
	go func() {
		router := NewRouter(s.cfg)

		n := negroni.Classic()
		n.UseHandler(router)

		addr := fmt.Sprintf(":%d", s.cfg.AppPort)
		log.Printf("✅ Servidor HTTP escuchando en http://localhost%s", addr)
		if err := http.ListenAndServe(addr, n); err != nil {
			log.Fatalf("❌ Error iniciando servidor HTTP: %v", err)
		}
	}()
}
