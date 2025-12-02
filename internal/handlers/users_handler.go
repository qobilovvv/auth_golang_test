package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/qobilovvv/test_tasks/auth/internal/services"
	"github.com/qobilovvv/test_tasks/auth/pkg/helpers"
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
		helpers.ResponseError(w, http.StatusBadRequest, "Invalid credentials")
		return
	}

	user_id, err := h.service.SignUpUser(req.OtpConfirmationToken, req.Email, req.Name, req.Password)
	if err != nil {
		helpers.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.RespondJSON(w, http.StatusOK, user_id)
}

func (h *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Identifier string `json:"phone_or_email"`
		Password   string `json:"password"`
		UserType   string `json:"user_type"` // only sysuser and user
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.ResponseError(w, http.StatusBadRequest, "Invalid credentials")
		return
	}

	if req.UserType == "" {
		req.UserType = "user"
	}

	token, err := h.service.Login(req.Identifier, req.Password, req.UserType)
	if err != nil {
		helpers.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.RespondJSON(w, http.StatusAccepted, map[string]string{"message": token})
}
