package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/urfave/negroni"
	"github.com/vboluda/go-backend-task-management/config"
	"github.com/vboluda/go-backend-task-management/database"
)

type Server struct {
	cfg *config.Config
	db  *database.Database
}

func NewServer(cfg *config.Config, db *database.Database) *Server {
	s := new(Server)
	s.cfg = cfg
	s.db = db
	return s
}

func (s *Server) Start() {
	go func() {
		router := NewRouter(s.cfg, s.db)

		n := negroni.Classic()
		n.UseHandler(router)

		addr := fmt.Sprintf(":%d", s.cfg.AppPort)
		log.Printf("✅ Servidor HTTP escuchando en http://localhost%s", addr)
		if err := http.ListenAndServe(addr, n); err != nil {
			log.Fatalf("❌ Error iniciando servidor HTTP: %v", err)
		}
	}()
}
