package config

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("secrett")

func GenerateOtpToken(otpID string, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"id":  otpID,
		"exp": time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func DecodeOtpToken(tokenString string) (string, time.Time, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtSecret, nil
	})
	if err != nil {
		return "", time.Time{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", time.Time{}, fmt.Errorf("invalid token")
	}

	otpIDRaw, ok := claims["id"]
	if !ok {
		return "", time.Time{}, fmt.Errorf("otp id missing")
	}
	otpID, ok := otpIDRaw.(string)
	if !ok {
		return "", time.Time{}, fmt.Errorf("otp id is not string")
	}

	expRaw, ok := claims["exp"]
	if !ok {
		return "", time.Time{}, fmt.Errorf("exp missing")
	}

	var expTime time.Time
	switch v := expRaw.(type) {
	case float64:
		expTime = time.Unix(int64(v), 0)
	case int64:
		expTime = time.Unix(v, 0)
	default:
		return "", time.Time{}, fmt.Errorf("invalid exp type")
	}

	return otpID, expTime, nil
}
