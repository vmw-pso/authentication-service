package main

import (
	"net/http"

	"github.com/vmw-pso/authentication-service/data"
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
		w.Write([]byte("Signed in successfully"))
	}
}
