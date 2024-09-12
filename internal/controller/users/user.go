package users

import (
	"github.com/gin-gonic/gin"
	"medodsAuth/internal/models"
	"net/http"
)

func Register(c *gin.Context) {
	var user *models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
}

func Login(c *gin.Context) {

}
