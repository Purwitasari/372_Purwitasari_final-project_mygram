package controllers

import (
	"MyGram/database"
	"MyGram/helpers"
	"MyGram/models"

	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type DataPhoto struct {
	ID        uint      `json:"photo_id"`
	Title     string    `json:"photo_title"`
	Caption   string    `json:"photo_caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    uint      `json:"photo_UserId"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
	User      DataUser
}

type DataUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

func PhotoCreate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	userID := uint(userData["id"].(float64))
	Photo := models.Photo{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"user_id":    Photo.UserID,
		"created_at": Photo.CreatedAt,
	})
}

func PhotoGet(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := []models.Photo{}
	User := models.User{}

	userID := uint(userData["id"].(float64))
	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err := db.Where("user_id = ?", userID).Find(&Photo).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": err.Error(),
		})
		return
	}

	err = db.Select("username,email").Where("id = ?", userID).Find(&User).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "ID Not Found",
		})
		return
	}
	user := DataUser{
		Email:    User.Email,
		Username: User.Username,
	}
	var data []DataPhoto
	for i := 0; i < len(Photo); i++ {
		temp := DataPhoto{
			ID:        Photo[i].ID,
			Title:     Photo[i].Title,
			Caption:   Photo[i].Caption,
			PhotoUrl:  Photo[i].PhotoUrl,
			UserID:    Photo[i].UserID,
			CreatedAt: Photo[i].CreatedAt,
			UpdatedAt: Photo[i].UpdatedAt,
			User:      user,
		}
		data = append(data, temp)
	}
	c.JSON(http.StatusOK, data)
}

func PhotoUpdate(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err := db.Model(&Photo).Where("id = ?", photoId).Updates(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"user_id":    Photo.UserID,
		"updated_at": Photo.UpdatedAt,
	})
}

func PhotoDelete(c *gin.Context) {
	db := database.GetDB()
	Photo := models.Photo{}
	photoID := c.Param("photoId")

	err := db.Model(&Photo).Where("id = ?", photoID).Delete(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Bad Request",
			"Message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your Photo Has Been Deleted",
	})
}
