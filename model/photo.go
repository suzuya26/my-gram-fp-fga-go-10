package model

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	Id       uint   `gorm:"primaryKey" json:"id"`
	Title    string `gorm:"not null" json:"title" form:"title" valid:"required~Judul Foto Wajib diisi"`
	Caption  string `gorm:"not null" json:"caption" form:"caption" valid:"required~Caption untuk Foto Wajib diisi"`
	PhotoUrl string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Link Photo Kamu tidak boleh kosong"`
	UserId   uint
	GormModel
	User *User
}

// hooks
func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}

	return
}
