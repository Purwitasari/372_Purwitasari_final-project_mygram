package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

type SocialMedia struct {
	gorm.Model
	Name      string `json:"sosmed_name" gorm:"not null" form:"sosmed_name" valid:"required"`
	SosMedUrl string `json:"sosmed_url" gorm:"not null" form:"sosmed_url" valid:"required"`
	UserID    uint   `json:"UserId" gorm:"index"`
	User      []User `gorm:"foreignKey:UserId" json:"user"`
}

func (sm *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(sm)

	if errCreate != nil {
		err = errCreate
		return
	}
	return
}
