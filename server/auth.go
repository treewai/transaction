package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/treewai/transaction/models"
)

func generateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	signedToken, err := token.SignedString([]byte("secret_key"))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (s *Server) AuthHandler(w http.ResponseWriter, r *http.Request) {
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !s.db.IsValidUser(&u) {
		token, err := generateToken(u.Username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(token))
		return
	}

	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}
