package controller

import (
	"my-gram/database"
	"my-gram/helper"
	"my-gram/model"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Createphoto(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)

	photo := model.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&photo)
	} else {
		c.ShouldBind(&photo)
	}

	photo.UserId = userID

	err := db.Debug().Create(&photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, photo)
}

func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)

	photoId, _ := strconv.Atoi(c.Param("photoId"))

	photo := model.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&photo)
	} else {
		c.ShouldBind(&photo)
	}

	photo.UserId = userID
	photo.Id = uint(photoId)

	err := db.Model(&photo).Where("id = ?", photoId).Updates(model.Photo{
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: photo.PhotoUrl,
	}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, photo)
}

func GetPhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	photoId, _ := strconv.Atoi(c.Param("photoId"))

	photo := model.Photo{}
	userID := uint(userData["id"].(float64))

	photo.UserId = userID
	photo.Id = uint(photoId)

	err := db.Preload("User", func(db_s *gorm.DB) *gorm.DB {
		return db_s.Select("id, user_name, age")
	}).First(&photo, photoId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})

		return
	}
	c.JSON(http.StatusOK, photo)
}

func GetAllPhoto(c *gin.Context) {
	db := database.GetDB()
	photo := []model.Photo{}

	err := db.Preload("User", func(db_s *gorm.DB) *gorm.DB {
		return db_s.Select("id, user_name, age")
	}).Find(&photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})

		return
	}
	c.JSON(http.StatusOK, photo)
}

func DeleteProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	photo := model.Photo{}

	photo.UserId = userID
	photo.Id = uint(photoId)

	err := db.Delete(&photo, photoId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})

		return
	}
	c.JSON(http.StatusNoContent, nil)
}
