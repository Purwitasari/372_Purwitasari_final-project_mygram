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

type DataComment struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoID   uint      `json:"photo_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      UserComment
	Photo     PhotoComment
}
type UserComment struct {
	ID       uint   `json:"id"`
	Email    string `json:"user_email"`
	Username string `json:"user_name"`
}

type PhotoComment struct {
	ID       uint   `json:"id"`
	Title    string `json:"photo_title"`
	Caption  string `json:"photo_caption"`
	PhotoUrl string `json:"photo_url"`
	UserID   uint   `json:"user_id"`
}

func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}
	photoId, _ := strconv.Atoi(c.Request.FormValue("photo_id"))

	Comment.UserID = userID
	Comment.PhotoID = uint(photoId)

	err := db.Debug().Create(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoID,
		"user_id":    Comment.UserID,
		"created_at": Comment.CreatedAt,
	})
}

func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	err := db.Model(&Comment).Where("id = ?", commentId).Updates(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	Comment.UserID = userID
	Comment.ID = uint(commentId)
	c.JSON(http.StatusOK, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoID,
		"user_id":    Comment.UserID,
		"updated_at": Comment.UpdatedAt,
	})
}

func GetComment(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	User := []models.User{}
	Comment := []models.Comment{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}
	type dataU struct {
		ID       uint   `json:"user_id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}
	var UserData []UserComment
	err := db.Find(&User).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": err.Error(),
		})
		return
	} else {

		for i := 0; i < len(User); i++ {
			StructData := UserComment{
				ID:       User[i].ID,
				Username: User[i].Username,
				Email:    User[i].Email,
			}
			UserData = append(UserData, StructData)
		}

	}

	err = db.Model(&Comment).Where("photo_id = ?", photoId).Find(&Comment).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": err.Error(),
		})
		return
	} else {
		var data []DataComment
		for i, v := range Comment {
			_ = v

			err = db.Model(&Photo).Where("id = ?", Comment[i].PhotoID).Find(&Photo).Error
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"error":   "Not Found Photo",
					"message": err.Error(),
				})
				return
			} else {
				var users UserComment
				for j := 0; j < len(User); j++ {
					if Comment[i].UserID == UserData[j].ID {
						users = UserComment{
							ID:       UserData[j].ID,
							Username: UserData[j].Username,
							Email:    UserData[j].Email,
						}
					}
				}
				photos := PhotoComment{
					ID:       Photo.ID,
					Title:    Photo.Title,
					Caption:  Photo.Caption,
					PhotoUrl: Photo.PhotoUrl,
					UserID:   Photo.UserID,
				}
				dataStruct := DataComment{
					ID:        Comment[i].ID,
					Message:   Comment[i].Message,
					PhotoID:   Comment[i].PhotoID,
					UserID:    Comment[i].UserID,
					UpdatedAt: Comment[i].UpdatedAt,
					CreatedAt: Comment[i].CreatedAt,
					User:      users,
					Photo:     photos,
				}
				data = append(data, dataStruct)
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}

func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	Comment := models.Comment{}
	commentId, _ := strconv.Atoi(c.Param("commentId"))

	err := db.Model(&Comment).Where("id = ?", commentId).Delete(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your Comment Has Been Deleted",
	})
}
