package tokens

import (
	"github.com/gin-gonic/gin"
	"medodsAuth/internal/models"
	storage "medodsAuth/internal/storage/postgresql"
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

// TODO RefreshTokens handler
func RefreshTokens(c *gin.Context) {

}
