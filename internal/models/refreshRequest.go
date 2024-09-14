package models

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

func (req *RefreshRequest) ParseAccessToken() (jwt.MapClaims, error) {
	const op = "models.refreshRequest.ParseAccessToken"
	secretKey := "NoSecretFromMedods"

	token, err := jwt.Parse(req.AccessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%s : %s", op, "signing problem")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("%s : %s", op, err.Error())
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("%s: invalid token", op)
}
