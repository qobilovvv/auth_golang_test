package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
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

	RespondJSON(w, http.StatusOK, map[string]string{"otp_id": otp.Id.String()})
}

func (h *otpHandler) ConfirmOTP(w http.ResponseWriter, r *http.Request) {
	var req struct {
		OtpId string `json:"otp_id"`
		Code  string `json:"code"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := uuid.Parse(req.OtpId)
	if err != nil {
		http.Error(w, "invalid otp_id", http.StatusBadRequest)
		return
	}

	jwtToken, err := h.service.ConfirmOTP(id, req.Code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	RespondJSON(w, http.StatusOK, map[string]string{"message": jwtToken})
}
