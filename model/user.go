package model

import (
	"my-gram/helper"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	Id       uint   `gorm:"primaryKey" json:"id"`
	UserName string `gorm:"not null;uniqueIndex" json:"user_name" form:"user_name" valid:"required~Nama Pengguna Kamu Wajib diisi"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email Kamu Wajib diisi"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~Password Wajib diisi,stringlength(6|999)~Password minimal 6 karakter"`
	Age      int    `gorm:"not null" json:"age" form:"age" valid:"required~Umur Wajib diisi,range(9|150)~Umur minimal 9 tahun"`
	GormModel
	SosialMedia []SosialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"social_media"`
	Photo       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
	Comment     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
}

// hooks
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	u.Password = helper.HashPass(u.Password)

	return
}

type UserReqEq struct {
	UserName string `json:"user_name" example:"budianduk"`
	Email    string `json:"email" example:"budi@sep.com"`
	Password string `json:"password" example:"budi69"`
	Age      int    `json:"age" example:"22"`
}

type LoginReq struct {
	Email    string `json:"email" example:"budi@sep.com"`
	Password string `json:"password" example:"budi69"`
}
