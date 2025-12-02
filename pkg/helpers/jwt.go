package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateJwtOtpToken(otpID string, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"id":  otpID,
		"exp": time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func DecodeJwtOtpToken(tokenString string) (string, time.Time, error) {
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

func GenerateAccessToken(user_id, user_type string, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id":   user_id,
		"user_type": user_type,
		"exp":       time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func DecodeAccessToken(tokenString string) (string, string, error) {
	t, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return "", "", err
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return "", "", fmt.Errorf("invalid token")
	}

	userID, _ := claims["user_id"].(string)
	userType, _ := claims["user_type"].(string)

	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			return "", "", fmt.Errorf("token expired")
		}
	}

	if userID == "" || userType == "" {
		return "", "", fmt.Errorf("missing fields")
	}

	return userID, userType, nil
}
