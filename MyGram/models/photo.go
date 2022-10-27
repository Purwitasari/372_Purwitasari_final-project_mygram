package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

type Photo struct {
	gorm.Model
	Title    string `json:"photo_title" gorm:"not null" form:"photo_title" valid:"required"`
	Caption  string `json:"photo_caption" gorm:"not null" form:"photo_caption" valid:"required"`
	PhotoUrl string `json:"photo_url" gorm:"not null" form:"photo_url" valid:"required"`
	UserID   uint   `json:"photo_UserId" gorm:"index"`
	User     []User `json:"user" gorm:"foreignKey:photo_UserId"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}
	return
}
