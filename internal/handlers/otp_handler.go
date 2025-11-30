package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/qobilovvv/test_tasks/auth/internal/services"
)

type otpHandler struct {
	service services.OTPService
}

func NewOTPHandler(service services.OTPService) *otpHandler {
	return &otpHandler{service: service}
}

func (h *otpHandler) SendOTP(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Email == "" {
		RespondJSON(w, http.StatusBadRequest, "email is required")
		return
	}

	otp, err := h.service.SendOTP(req.Email)
	if err != nil {
		RespondJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"otp_id": otp.Id.String(),
	})
}