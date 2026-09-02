package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"bigoproject/domain"
	"bigoproject/usecase"
)

type CreateUserRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type UserHandler struct {
	registerUser *usecase.RegisterUserUseCase
}

func NewUserHandler(uc *usecase.RegisterUserUseCase) *UserHandler {
	return &UserHandler{registerUser: uc}
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}

	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	cmd := usecase.RegisterUserCommand{
		Email: req.Email,
		Name:  req.Name,
	}

	user, err := h.registerUser.Execute(r.Context(), cmd)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrInvalidEmail):
			http.Error(w, "invalid email", http.StatusBadRequest)
		case errors.Is(err, domain.ErrEserAlreadyExists):
			http.Error(w, "user already exists", http.StatusConflict)
		default:
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}