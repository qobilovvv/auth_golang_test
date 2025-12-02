package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/qobilovvv/test_tasks/auth/internal/services"
	"github.com/qobilovvv/test_tasks/auth/pkg/helpers"
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
		helpers.RespondJSON(w, http.StatusBadRequest, "email is required")
		return
	}

	otp, err := h.service.SendOTP(req.Email)
	if err != nil {
		helpers.RespondJSON(w, http.StatusInternalServerError, "something went wrong")
		log.Println("Error in handler:", err)
		return
	}

	helpers.RespondJSON(w, http.StatusOK, map[string]string{"otp_id": otp.Id.String()})
}

func (h *otpHandler) ConfirmOTP(w http.ResponseWriter, r *http.Request) {
	var req struct {
		OtpId string `json:"otp_id"`
		Code  string `json:"code"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.ResponseError(w, http.StatusBadRequest, "something went wrong")
		log.Println("Error in handler:", err)
		return
	}

	id, err := uuid.Parse(req.OtpId)
	if err != nil {
		helpers.ResponseError(w, http.StatusBadRequest, "invalid otp_id")
		return
	}

	jwtToken, err := h.service.ConfirmOTP(id, req.Code)
	if err != nil {
		helpers.ResponseError(w, http.StatusInternalServerError, "something went wrong")
		log.Println("Error in handler:", err)
		return
	}

	helpers.RespondJSON(w, http.StatusOK, map[string]string{"message": jwtToken})
}
