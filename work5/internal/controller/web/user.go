package web

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/mch735/education/work5/internal/entity"
	"github.com/mch735/education/work5/internal/entity/dto"
	"github.com/mch735/education/work5/internal/usecase"
)

func Index(uc usecase.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := uc.GetUsers(r.Context())
		if err != nil {
			slog.Error(fmt.Sprintf("uc.GetUsers: %v", err))
			errorResponse(w, "database problems", http.StatusInternalServerError)

			return
		}

		jsonResponse(w, users)
	}
}

func Show(uc usecase.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := uc.GetUserByID(r.Context(), r.PathValue("id"))
		if err != nil {
			slog.Error(fmt.Sprintf("uc.GetUserByID: %v", err))
			errorResponse(w, "database problems", http.StatusInternalServerError)

			return
		}

		jsonResponse(w, user)
	}
}

func Create(uc usecase.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := dto.User{}

		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			slog.Error(fmt.Sprintf("json.NewDecoder: %v", err))
			errorResponse(w, "invalid request body", http.StatusBadRequest)

			return
		}

		user := entity.User{
			Name:  u.Name,
			Email: u.Email,
			Role:  u.Role,
		}

		err = uc.Create(r.Context(), &user)
		if err != nil {
			slog.Error(fmt.Sprintf("uc.Create: %v", err))
			errorResponse(w, "database problems", http.StatusInternalServerError)

			return
		}

		jsonResponse(w, user)
	}
}

func Update(uc usecase.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := dto.User{
			ID: r.PathValue("id"),
		}

		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			slog.Error(fmt.Sprintf("json.NewDecoder: %v", err))
			errorResponse(w, "invalid request body", http.StatusBadRequest)

			return
		}

		user := entity.User{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
			Role:  u.Role,
		}

		err = uc.Update(r.Context(), &user)
		if err != nil {
			slog.Error(fmt.Sprintf("uc.Update: %v", err))
			errorResponse(w, "database problems", http.StatusInternalServerError)

			return
		}

		jsonResponse(w, user)
	}
}

func Delete(uc usecase.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := uc.Delete(r.Context(), r.PathValue("id"))
		if err != nil {
			slog.Error(fmt.Sprintf("uc.Delete: %v", err))
			errorResponse(w, "database problems", http.StatusInternalServerError)

			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
