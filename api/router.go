package api

import (
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/vboluda/go-backend-task-management/config"
	_ "github.com/vboluda/go-backend-task-management/docs" // Swagger docs auto-generadas
)

func NewRouter(cfg *config.Config) http.Handler {
	r := mux.NewRouter()

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Subrouter para /api/user
	userHandler := NewUserHandler(cfg)
	userRouter := r.PathPrefix("/api/user").Subrouter()
	userRouter.HandleFunc("/login", userHandler.Login).Methods("POST")
	userRouter.HandleFunc("/logout", userHandler.Logout).Methods("POST")
	userRouter.HandleFunc("/change-password", userHandler.ChangePassword).Methods("POST")

	return r
}
