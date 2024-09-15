package users

import (
	"github.com/gin-gonic/gin"
	"log"
	"medodsAuth/internal/models"
	storage "medodsAuth/internal/storage/postgresql"
	"net/http"
)

func Register(c *gin.Context) {
	var user *models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := user.CreateUser(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := storage.DB.SaveUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("user has created with guid: %s", user.GUID)

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"guid":    user.GUID,
	})
}
