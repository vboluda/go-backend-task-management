package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/vboluda/go-backend-task-management/api"
	"github.com/vboluda/go-backend-task-management/config"
)

func main() {
	// 🔧 Cargar configuración desde .env o entorno
	cfg := config.New().LoadEnv().Validate()

	// 📋 Mostrar configuración por consola
	fmt.Println(cfg)

	server := api.NewServer(cfg)
	server.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("🕐 Servidor iniciado. Presioná Ctrl+C para salir...")
	<-stop

	fmt.Println("\n👋 Señal recibida. Cerrando aplicación.")
	// Aquí podrías cerrar conexiones, limpiar recursos, etc.
}
