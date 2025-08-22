package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/vboluda/go-backend-task-management/config"
)

type Database struct {
	DB *sql.DB
}

// Init abre la conexión y prepara la base de datos
func Init(cfg *config.Config) *Database {
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("❌ No se pudo conectar a la base de datos: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("❌ La base de datos no respondió: %v", err)
	}

	log.Println("✅ Conectado a PostgreSQL")

	d := &Database{DB: db}
	d.runSchema()

	return d
}
