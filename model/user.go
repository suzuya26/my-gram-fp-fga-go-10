package model

import (
	"my-gram/helper"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	Id       int    `gorm:"primaryKey;type:serial" json:"id"`
	FullName string `gorm:"not null;uniqueIndex" json:"full_name" form:"full_name" valid:"required~Nama Lengkap Kamu Wajib diisi"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email Kamu Wajib diisi"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~Password Wajib diisi"`
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
