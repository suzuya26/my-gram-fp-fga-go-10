package model

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SosialMedia struct {
	Id             int    `gorm:"primaryKey;type:serial" json:"id"`
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required~Nama Sosial Media kamu tidak boleh kosong"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~Link Sosial Media Kamu tidak boleh kosong"`
	UserId         uint
	GormModel
	User *User
}

// hooks
func (sm *SosialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(sm)
	if err != nil {
		return err
	}

	return
}
