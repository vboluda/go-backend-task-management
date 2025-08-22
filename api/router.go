package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vboluda/go-backend-task-management/config"
)

func NewRouter(cfg *config.Config) http.Handler {
	r := mux.NewRouter()

	// Subrouter para /api/user
	userHandler := NewUserHandler(cfg)
	userRouter := r.PathPrefix("/api/user").Subrouter()

	userRouter.HandleFunc("/login", userHandler.Login).Methods("POST")
	userRouter.HandleFunc("/logout", userHandler.Logout).Methods("POST")
	userRouter.HandleFunc("/change-password", userHandler.ChangePassword).Methods("POST")

	return r
}
