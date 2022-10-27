package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Message string  `json:"message" gorm:"not null" form:"message" valid:"required"`
	UserID  uint    `json:"user_id" gorm:"index"`
	PhotoID uint    `json:"photo_id" gorm:"index"`
	Photo   []Photo `json:"photo" gorm:"foreignKey:photo_id"`
	User    []User  `json:"user" gorm:"foreignKey:user_id"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}
	return
}
