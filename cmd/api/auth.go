package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/vmw-pso/authentication-service/data"
	"golang.org/x/crypto/bcrypt"
)

var signingKey string = "supersecretkey" // here for dev, until drawn from the store as the correct key

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func validatePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateJWT(user data.User) (string, error) {
	// var signingKey = []byte(os.Getenv("tokenkey"))

	token := jwt.New(jwt.SigningMethodEdDSA)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(24 * time.Hour)
	claims["authorized"] = true
	claims["username"] = user.Username
	claims["email"] = user.Email

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *server) verifyToken(endpointHander func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] == nil {
			s.tools.ErrorJSON(w, errors.New("Unauthorized"), http.StatusUnauthorized)
			return
		}

		// CONTINUE HERE!!!
	}
}
