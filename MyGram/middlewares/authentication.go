package middlewares

import (
	"MyGram/database"
	"MyGram/helpers"
	"MyGram/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": err.Error(),
			})
			return
		}

		c.Set("userData", verifyToken)
		c.Next()
	}
}

func UserAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		UserID := uint(userData["id"].(float64))
		User := models.User{}

		err := db.Where("id = ?", UserID).First(&User).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "Data Doest'n Exist",
			})
			return
		}
		if User.ID != UserID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You Are Not Allowed To Access This Data",
			})
		}
		c.Next()
	}
}
