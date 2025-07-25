package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	row := db.QueryRow("SELECT user_id, full_name FROM users WHERE email = $1", user.Email)
	err := row.Scan(&user.ID, &user.FullName)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, _ := token.SignedString(jwtKey)

	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}

func CreateTerminalHandler(w http.ResponseWriter, r *http.Request) {
	var terminal Terminal
	_ = json.NewDecoder(r.Body).Decode(&terminal)

	err := db.QueryRow("INSERT INTO terminals (terminal_name) VALUES ($1) RETURNING terminal_id", terminal.Name).Scan(&terminal.ID)
	if err != nil {
		http.Error(w, "Insert failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(terminal)
}
