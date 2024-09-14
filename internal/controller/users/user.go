package users

import (
	"github.com/gin-gonic/gin"
	"medodsAuth/internal/models"
	storage "medodsAuth/internal/storage/postgresql"
	"net/http"
)

func Register(c *gin.Context) {
	var user *models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if err := user.CreateUser(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := storage.DB.SaveUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"guid":    user.GUID,
	})
}
