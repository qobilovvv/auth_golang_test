package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/qobilovvv/test_tasks/auth/internal/services"
)

type userHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *userHandler {
	return &userHandler{service: service}
}

func (h *userHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req struct {
		OtpConfirmationToken string `json:"otp_confirmation_token"`
		Email                string `json:"email"`
		Name                 string `json:"name"`
		Password             string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponseError(w, http.StatusBadRequest, "Invalid credentials")
		return
	}

	user_id, err := h.service.SignUpUser(req.OtpConfirmationToken, req.Email, req.Name, req.Password)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, user_id)
}

func (h *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		UserType string `json:"user_type"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponseError(w, http.StatusBadRequest, "Invalid credentials")
		return
	}

	if req.UserType == "" {
		req.UserType = "user"
	}

	token, err := h.service.Login(req.Email, req.Password, req.UserType)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondJSON(w, http.StatusAccepted, map[string]string{"message": token})
}
