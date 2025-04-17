package web

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/mch735/education/work5/internal/entity"
	"github.com/mch735/education/work5/internal/entity/dto"
	"github.com/mch735/education/work5/internal/usecase"
)

func Index(uc usecase.User) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		users, err := uc.GetUsers()
		if err != nil {
			slog.Error("web - v1 - index", slog.String("error", err.Error()))
			errorResponse(w, "database problems", http.StatusInternalServerError)

			return
		}

		jsonResponse(w, users)
	}
}

func Show(uc usecase.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := uc.GetUserByID(r.PathValue("id"))
		if err != nil {
			slog.Error("web - v1 - show", slog.String("error", err.Error()))
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
			slog.Error("web - v1 - create", slog.String("error", err.Error()))
			errorResponse(w, "invalid request body", http.StatusBadRequest)

			return
		}

		user := entity.User{
			Name:  u.Name,
			Email: u.Email,
			Role:  u.Role,
		}

		err = uc.Create(&user)
		if err != nil {
			slog.Error("web - v1 - create", slog.String("error", err.Error()))
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
			slog.Error("web - v1 - update", slog.String("error", err.Error()))
			errorResponse(w, "invalid request body", http.StatusBadRequest)

			return
		}

		user := entity.User{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
			Role:  u.Role,
		}

		err = uc.Update(&user)
		if err != nil {
			slog.Error("web - v1 - update", slog.String("error", err.Error()))
			errorResponse(w, "database problems", http.StatusInternalServerError)

			return
		}

		jsonResponse(w, user)
	}
}

func Delete(uc usecase.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := uc.Delete(r.PathValue("id"))
		if err != nil {
			slog.Error("web - v1 - delete", slog.String("error", err.Error()))
			errorResponse(w, "database problems", http.StatusInternalServerError)

			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
