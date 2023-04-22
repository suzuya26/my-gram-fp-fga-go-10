package model

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	Id      uint `gorm:"primaryKey" json:"id"`
	UserId  uint
	PhotoId int    `json:"photo_id" form:"photo_id" valid:"required~Id Photo yang dikomentari tidak boleh kosong"`
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~Pesan Komentar tidak boleh kosong"`
	GormModel
	User  *User
	Photo *Photo
}

type ErrorNotFoundResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (cm *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(cm)
	if err != nil {
		return err
	}

	//check exist photo yg dikomen
	photoId := cm.PhotoId

	var photo Photo
	if err := tx.Where("id =?", photoId).First(&photo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("photo Id tidak valid")
		}
		return err
	}
	return
}

type CommentReqEq struct {
	PhotoId int    `json:"photo_id" example:"1"`
	Message string `gorm:"not null" json:"message" example:"Ambis Sekali"`
}

type UpdateCommentReq struct {
	PhotoId int    `json:"photo_id" example:"1"`
	Message string `gorm:"not null" json:"message" example:"Ambis Sekali yah bund"`
}
