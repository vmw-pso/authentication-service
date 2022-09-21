package main

import (
	"net/http"

	"github.com/vmw-pso/authentication-service/data"
	"github.com/vmw-pso/toolkit"
)

func (s *server) handleSignup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestPayload struct {
			Name     string `json:"name"`
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
			Name:         requestPayload.Name,
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
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := s.tools.ReadJSON(w, r, &requestPayload); err != nil {
			s.tools.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		user, err := s.models.User.GetByEmail(s.DB, requestPayload.Email)
		if err != nil {
			s.tools.ErrorJSON(w, err, http.StatusUnauthorized)
			return
		}

		match, err := user.PasswordMatches(requestPayload.Password)
		if err != nil || !match {
			_ = s.tools.ErrorJSON(w, err, http.StatusUnauthorized)
			return
		}

		payload := toolkit.JSONResponse{
			Error:   false,
			Message: "Signed in",
			Data:    user,
		}

		_ = s.tools.WriteJSON(w, http.StatusAccepted, payload)
	}
}
