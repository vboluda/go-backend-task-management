package database

import (
	"database/sql"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// CreateUser inserta un nuevo usuario con la contraseña hasheada
func (d *Database) CreateUser(username, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = d.DB.Exec(`INSERT INTO users (username, password) VALUES ($1, $2)`, username, string(hashed))
	if err != nil {
		return err
	}

	log.Printf("✅ Usuario '%s' creado correctamente", username)
	return nil
}

// AuthenticateUser verifica el usuario y la contraseña
func (d *Database) AuthenticateUser(username, password string) (bool, error) {
	var hashed string

	err := d.DB.QueryRow(`SELECT password FROM users WHERE username = $1`, username).Scan(&hashed)
	if errors.Is(err, sql.ErrNoRows) {
		log.Println("⚠️ Usuario no encontrado")
		return false, nil // o manejo custom
	}

	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		return false, nil // contraseña incorrecta
	}

	return true, nil
}

// ChangePassword actualiza la contraseña del usuario
func (d *Database) ChangePassword(username, newPassword string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	result, err := d.DB.Exec(`UPDATE users SET password = $1 WHERE username = $2`, string(hashed), username)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("⚠️ usuario no encontrado")
	}

	log.Printf("🔐 Contraseña actualizada para usuario '%s'", username)
	return nil
}
