package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(secretKey []byte, userID string, userType string) (*string, *string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims (data dalam token)
	expiredTime := time.Now().Add(time.Hour * 24).Unix()
	expiredTimeISO := time.Unix(expiredTime, 0).Format(time.RFC3339)
	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"] = userID
	claims["user_type"] = userType
	claims["exp"] = expiredTime

	// Generate signed token string
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, nil, err
	}

	return &tokenString, &expiredTimeISO, nil
}
