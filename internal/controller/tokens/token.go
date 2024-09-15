package tokens

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"medodsAuth/internal/models"
	storage "medodsAuth/internal/storage/postgresql"
	"medodsAuth/internal/utils"
	"net/http"
)

func GetTokens(c *gin.Context) {
	var user *models.User
	guid := c.Query("guid")
	user, err := storage.DB.GetUserByGUID(guid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	ip := c.ClientIP()

	tokens, err := user.GenerateTokens(ip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	refreshTokenInfo := &models.RefreshToken{
		RefreshTokenHash: tokens.RefreshTokenHash,
		GUID:             guid,
		IP:               ip,
		ID:               tokens.TokenID,
	}
	if err := storage.DB.SaveToken(refreshTokenInfo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  tokens.Access,
		"refresh_token": tokens.Refresh,
	})
}

func RefreshTokens(c *gin.Context) {
	const (
		op = "controller.tokens.RefreshTokens"
	)
	var request *models.RefreshRequest
	if err := c.BindJSON(&request); err != nil {
		log.Printf("%s : %s", op, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	claims, err := request.ParseAccessToken()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token"})
		return
	}

	guid := claims["GUID"].(string)
	tokenID := claims["TokenID"].(string)
	ip := claims["IPAddress"].(string)

	storedRefreshToken, err := storage.DB.GetRefreshToken(guid, tokenID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedRefreshToken.RefreshTokenHash), []byte(request.RefreshToken)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	user, err := storage.DB.GetUserByGUID(guid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	currentIP := c.ClientIP()
	if currentIP != ip {
		utils.SendEmailWarning(user.Email)
	}

	newTokens, err := user.GenerateTokens(currentIP)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	if err := storage.DB.UpdateRefreshToken(guid, tokenID, newTokens.RefreshTokenHash); err != nil {
		log.Printf("%s : %s", op, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update tokens"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  newTokens.Access,
		"refresh_token": newTokens.Refresh,
	})
}
