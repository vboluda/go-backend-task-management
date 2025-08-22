package api

import (
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/vboluda/go-backend-task-management/config"
	"github.com/vboluda/go-backend-task-management/database"
	_ "github.com/vboluda/go-backend-task-management/docs" // Swagger docs auto-generadas
)

func NewRouter(cfg *config.Config, db *database.Database) http.Handler {
	r := mux.NewRouter()
	auth := AuthMiddleware(cfg)

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Subrouter para /api/user
	userHandler := NewUserHandler(cfg, db)
	userRouter := r.PathPrefix("/api/user").Subrouter()
	userRouter.HandleFunc("/login", userHandler.Login).Methods("POST")
	userRouter.HandleFunc("/logout", userHandler.Logout).Methods("POST").Handler(auth(http.HandlerFunc(userHandler.Logout)))
	userRouter.HandleFunc("/change-password", userHandler.ChangePassword).Methods("POST").Handler(auth(http.HandlerFunc(userHandler.ChangePassword)))

	return r
}
