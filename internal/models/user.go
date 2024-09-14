package models

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
	"time"
)

type User struct {
	GUID      string `json:"GUID"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (user *User) validate() error {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	user.Email = strings.TrimSpace(user.Email)
	if user.Email == "" {
		return fmt.Errorf("invalid email")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return fmt.Errorf("invalid password")
	}

	return nil
}

// CreateUser encrypts user password and validates a data
func (user *User) CreateUser() error {
	if err := user.validate(); err != nil {
		return err
	}

	pwSlice, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return fmt.Errorf("failed to encrypt the password")
	}
	user.Password = string(pwSlice)

	user.GUID = uuid.New().String()

	return nil
}

func (user *User) GenerateTokens(ipAddress string) (*Tokens, error) {
	const op = "models.users.GenerateTokens"

	tokenID := uuid.New().String()

	accessToken, err := generateAccessToken(user.GUID, ipAddress, tokenID)
	if err != nil {
		log.Printf("%s : %s\n", op, err.Error())
		return nil, fmt.Errorf("%s : %w", op, err)
	}

	refreshToken, refreshTokenHash, err := generateRefreshToken()
	if err != nil {
		log.Printf("%s : %s\n", op, err.Error())
		return nil, fmt.Errorf("%s : %w", op, err)
	}

	return &Tokens{
		Access:           accessToken,
		Refresh:          refreshToken,
		RefreshTokenHash: refreshTokenHash,
		TokenID:          tokenID,
	}, nil
}

func generateAccessToken(guid, ipAddress, tokenID string) (string, error) {
	claims := jwt.MapClaims{
		"GUID":      guid,
		"IPAddress": ipAddress,
		"TokenID":   tokenID,
		"exp":       time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	secretKey := "NoSecretForMedods"
	return token.SignedString([]byte(secretKey))
}

func generateRefreshToken() (string, string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", "", err
	}

	token := base64.StdEncoding.EncodeToString(b)

	hash, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}

	return token, string(hash), nil
}
