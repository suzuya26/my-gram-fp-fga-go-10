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

func CreateSosmed(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)

	sosmed := model.SosialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&sosmed)
	} else {
		c.ShouldBind(&sosmed)
	}

	sosmed.UserId = userID

	err := db.Debug().Create(&sosmed).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, sosmed)
}

func UpdateSosmed(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)

	sosmedId, _ := strconv.Atoi(c.Param("sosmedId"))

	sosmed := model.SosialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&sosmed)
	} else {
		c.ShouldBind(&sosmed)
	}

	sosmed.UserId = userID
	sosmed.Id = uint(sosmedId)

	err := db.Model(&sosmed).Where("id = ?", sosmedId).Updates(model.SosialMedia{
		Name:           sosmed.Name,
		SocialMediaUrl: sosmed.SocialMediaUrl,
	}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, sosmed)
}

func GetSosmed(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	sosmedId, _ := strconv.Atoi(c.Param("sosmedId"))

	sosmed := model.SosialMedia{}
	userID := uint(userData["id"].(float64))

	sosmed.UserId = userID
	sosmed.Id = uint(sosmedId)

	err := db.Preload("User", func(db_s *gorm.DB) *gorm.DB {
		return db_s.Select("id, user_name, age")
	}).First(&sosmed, sosmedId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})

		return
	}
	c.JSON(http.StatusOK, sosmed)
}

func GetAllSosmed(c *gin.Context) {
	db := database.GetDB()
	sosmed := []model.SosialMedia{}

	err := db.Preload("User", func(db_s *gorm.DB) *gorm.DB {
		return db_s.Select("id, user_name, age")
	}).Find(&sosmed).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})

		return
	}
	c.JSON(http.StatusOK, sosmed)
}

func DeleteSosmed(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	sosmedId, _ := strconv.Atoi(c.Param("sosmedId"))
	userID := uint(userData["id"].(float64))

	sosmed := model.SosialMedia{}

	sosmed.UserId = userID
	sosmed.Id = uint(sosmedId)

	err := db.Delete(&sosmed, sosmedId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})

		return
	}
	c.JSON(http.StatusNoContent, nil)
}
