package models

import (
	"MyGram/helpers"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"user_name" gorm:"not null,unique_index" form:"user_name" valid:"required"`
	Email    string `json:"user_email" gorm:"not null,unique_index" form:"user_email" valid:"email, required"`
	Password string `json:"user_password" gorm:"not null" form:"user_password" valid:"required"`
	Age      int    `json:"user_age" gorm:"not null" form:"user_age" valid:"required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}
	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
