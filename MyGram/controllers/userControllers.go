package controllers

import (
	"MyGram/database"
	"MyGram/helpers"
	"MyGram/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UserGet(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	User := models.User{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Where("id = ?", userID).Find(&User).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "ID Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       userID,
		"username": User.Username,
		"email":    User.Email,
		"age":      User.Age,
	})
}

func UserUpdate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	User := models.User{}

	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Model(&User).Where("id = ?", userID).Updates(&User).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         userID,
		"username":   User.Username,
		"email":      User.Email,
		"age":        User.Age,
		"updated_at": User.UpdatedAt,
	})
}

func UserDelete(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	User := models.User{}

	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Model(&User).Where("id = ?", userID).Delete(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your Account Has Been Deleted",
	})
}
