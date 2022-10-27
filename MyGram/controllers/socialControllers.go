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

type DataSosmed struct {
	ID        uint      `json:"id"`
	Name      string    `json:"sosmed_name"`
	SosMedUrl string    `json:"sosmed_url"`
	UserID    uint      `json:"UserId"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      UserSosmed
}

type UserSosmed struct {
	ID       uint   `json:"id"`
	Username string `json:"user_name"`
}

func CreateSosmed(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Sosmed := models.SocialMedia{}

	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Sosmed)
	} else {
		c.ShouldBind(&Sosmed)
	}

	Sosmed.UserID = userID

	err := db.Debug().Create(&Sosmed).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               Sosmed.ID,
		"name":             Sosmed.Name,
		"social_media_url": Sosmed.SosMedUrl,
		"user_id":          Sosmed.UserID,
		"created_at":       Sosmed.CreatedAt,
	})
}

func UpdateSosmed(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Sosmed := models.SocialMedia{}

	sosmedId, _ := strconv.Atoi(c.Param("sosmedId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Sosmed)
	} else {
		c.ShouldBind(&Sosmed)
	}

	Sosmed.UserID = userID
	Sosmed.ID = uint(sosmedId)

	err := db.Model(&Sosmed).Where("id = ?", sosmedId).Updates(&Sosmed).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               Sosmed.ID,
		"name":             Sosmed.Name,
		"social_media_url": Sosmed.SosMedUrl,
		"user_id":          Sosmed.UserID,
		"updated_at":       Sosmed.UpdatedAt,
	})
}

func GetSosmed(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Sosmed := []models.SocialMedia{}
	User := models.User{}

	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Sosmed)
	} else {
		c.ShouldBind(&Sosmed)
	}

	err := db.Where("user_id = ?", userID).Find(&Sosmed).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": err.Error(),
		})
		return
	}

	var data []DataSosmed
	for i := 0; i < len(Sosmed); i++ {
		err := db.Select("username, id").Where("id = ?", userID).Find(&User).Error
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": err.Error(),
			})
		}

		users := UserSosmed{
			ID:       uint(userID),
			Username: User.Username,
		}

		dataStruct := DataSosmed{
			ID:        Sosmed[i].ID,
			Name:      Sosmed[i].Name,
			SosMedUrl: Sosmed[i].SosMedUrl,
			UserID:    Sosmed[i].UserID,
			CreatedAt: Sosmed[i].CreatedAt,
			UpdatedAt: Sosmed[i].UpdatedAt,
			User:      users,
		}
		data = append(data, dataStruct)
	}
	c.JSON(http.StatusOK, gin.H{
		"sosmed": data,
	})
}

func DeleteSosmed(c *gin.Context) {
	db := database.GetDB()
	Sosmed := models.SocialMedia{}
	SosmedID := c.Param("socialMediaId")

	err := db.Model(&Sosmed).Where("id = ?", SosmedID).Delete(&Sosmed).Error
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
