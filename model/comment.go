package model

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	Id      int `gorm:"primaryKey;type:serial" json:"id"`
	UserId  uint
	PhotoId uint
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~Pesan Komentar tidak boleh kosong"`
	GormModel
	User  *User
	Photo *Photo
}

func (cm *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(cm)
	if err != nil {
		return err
	}

	return
}
