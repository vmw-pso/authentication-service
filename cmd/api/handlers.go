package main

import (
	"fmt"
	"net/http"

	"github.com/vmw-pso/authentication-service/data"
	"github.com/vmw-pso/toolkit"
)

func (s *server) handleSignup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestPayload struct {
			Username string `json:"username"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		err := s.tools.ReadJSON(w, r, &requestPayload)
		if err != nil {
			_ = s.tools.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		hash, err := hashPassword(requestPayload.Password)
		if err != nil {
			_ = s.tools.ErrorJSON(w, err, http.StatusInternalServerError)
			return
		}

		user := data.User{
			Username:     requestPayload.Username,
			Email:        requestPayload.Email,
			PasswordHash: hash,
		}

		id, err := user.Insert(s.DB)
		if err != nil {
			_ = s.tools.ErrorJSON(w, err, http.StatusInternalServerError)
			return
		}

		user.ID = id
		s.tools.WriteJSON(w, http.StatusAccepted, user)
	}
}

func (s *server) handleSignin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestPayload struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		fmt.Println(requestPayload.Username)

		if err := s.tools.ReadJSON(w, r, &requestPayload); err != nil {
			s.tools.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		user, err := s.models.User.GetByUsername(s.DB, requestPayload.Username)
		if err != nil {
			s.tools.ErrorJSON(w, err, http.StatusUnauthorized)
			return
		}

		match, err := user.PasswordMatches(requestPayload.Password)
		if err != nil || !match {
			_ = s.tools.ErrorJSON(w, err, http.StatusUnauthorized)
			return
		}

		jwt, err := generateJWT(*user)
		if err != nil {
			s.tools.ErrorJSON(w, err, http.StatusInternalServerError)
			return
		}

		payload := toolkit.JSONResponse{
			Error:   false,
			Message: "Signed in",
			Data:    jwt,
		}

		_ = s.tools.WriteJSON(w, http.StatusAccepted, payload)
	}
}
