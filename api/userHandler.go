package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/vboluda/go-backend-task-management/config"
	"github.com/vboluda/go-backend-task-management/database"
)

type UserHandler struct {
	cfg *config.Config
	db  *database.Database
}

func NewUserHandler(cfg *config.Config, db *database.Database) *UserHandler {
	return &UserHandler{cfg: cfg, db: db}
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
}

// @Summary Login de usuario
// @Description Genera un JWT válido para el usuario autenticado
// @Tags user
// @Accept json
// @Produce json
// @Param credentials body loginRequest true "Credenciales de usuario"
// @Success 200 {object} loginResponse
// @Failure 400 {string} string "Solicitud inválida"
// @Failure 500 {string} string "Error interno"
// @Router /api/user/login [post]
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}

	ok, err := h.db.AuthenticateUser(req.Email, req.Password)
	if err != nil {
		http.Error(w, "Error interno al autenticar", http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "Credenciales inválidas", http.StatusUnauthorized)
		return
	}

	// Crear el JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": req.Email,
		"exp":   time.Now().Add(1 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(h.cfg.JWTSecret))
	if err != nil {
		http.Error(w, "No se pudo generar el token", http.StatusInternalServerError)
		return
	}

	resp := loginResponse{Token: tokenString}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// @Summary Logout del usuario
// @Description Termina la sesión del usuario (simulado)
// @Tags user
// @Success 200 {string} string "Logout exitoso"
// @Router /api/user/logout [post]
// @Security BearerAuth
func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logout exitoso (simulado)"))
}

// @Summary Cambio de contraseña
// @Description Cambia la contraseña del usuario (simulado)
// @Tags user
// @Success 200 {string} string "Contraseña cambiada"
// @Router /api/user/change-password [post]
// @Security BearerAuth
func (h *UserHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Contraseña cambiada (simulado)"))
}
