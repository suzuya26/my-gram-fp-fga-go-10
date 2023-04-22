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

type PhotoReqEq struct {
	Title    string `json:"title" example:"ch-09 10"`
	Caption  string `json:"caption" example:"Salah satu SS challange-09"`
	PhotoUrl string `json:"photo_url" example:"https://res.cloudinary.com/drakr4vtu/image/upload/v1681832825/FGA%20GO%2010/challange-09/image_2023-04-18_224711118_xcyqhz.png"`
}

type UpdatePhotoReq struct {
	Title    string `json:"title" example:"ch-09 10-0"`
	Caption  string `json:"caption" example:"Salah satu SrenShot challange-09"`
	PhotoUrl string `json:"photo_url" example:"https://res.cloudinary.com/drakr4vtu/image/upload/v1681832825/FGA%20GO%2010/challange-09/image_2023-04-18_224711118_xcyqhz.png"`
}
