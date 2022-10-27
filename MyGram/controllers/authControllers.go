package controllers

import (
	"MyGram/database"
	"MyGram/helpers"
	"MyGram/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	passCheck := c.Request.FormValue("user_password")
	ageVar := c.Request.FormValue("user_age")
	age, err := strconv.Atoi(ageVar)
	_ = err

	errUsername := db.Model(&User).Where("username = ?", User.Username).First(&User).Error
	errEmail := db.Model(&User).Where("email = ?", User.Email).First(&User).Error
	if errUsername == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Username Already Exist",
		})
		return
	} else if errEmail == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Email Already Exist",
		})
		return
	}

	if len(passCheck) < 6 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Password Must Be At Least 6 Characters Long",
		})
		return
	} else {
		if age < 8 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Minimum Age 8 Years Old",
			})
			return
		} else {
			err := db.Debug().Create(&User).Error
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error":   "Bad Request",
					"message": err.Error(),
				})
				return
			}

			c.JSON(http.StatusCreated, gin.H{
				"id":       User.ID,
				"email":    User.Email,
				"username": User.Username,
				"age":      User.Age,
			})
		}
	}
}

func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unatuhorized",
			"message": "Invalid email",
		})
		return
	}
	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
