package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Router(serviceFactory func(ctx context.Context) Service) http.Handler {
	r := chi.NewRouter()

	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		// ✅ Now txFromContext() is called during a request
		service := serviceFactory(r.Context())

		user, err := service.Register(r.Context(), req.Name, req.Email, req.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp := struct {
			ID    string `json:"id"`
			Name  string `json:"name"`
			Email string `json:"email"`
		}{
			ID: user.ID, Name: user.Name, Email: user.Email,
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resp)
	})

	return r
}
