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

// CreateSosmed Create a new social media account
// @Summary Create a new social media account
// @Description Create a new social media account
// @Tags Social Media
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization header in the format of 'Bearer {token}'"
// @Param sosmed body model.SosmedReq true "Sosmed data"
// @Success 201 {object} helper.CreateSosmedResponse
// @Failure 400 {object} helper.ErrorResponse
// @Router /sosmed [post]
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

	c.JSON(http.StatusCreated, gin.H{
		"message":          "Success Added Social Media",
		"id":               sosmed.Id,
		"name":             sosmed.Name,
		"social_media_url": sosmed.SocialMediaUrl,
		"user_id":          sosmed.UserId,
	})
}

// UpdateSosmed Update a social media info
// @Summary Update existing sosmed
// @Description Update existing sosmed
// @Tags Social Media
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization header in the format of 'Bearer {token}'"
// @Param sosmedId path int true "Sosmed ID"
// @Param sosmed body model.UpdateSosmedReq true "Update sosmed object"
// @Success 200 {object} helper.UpdateSosmedResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 401 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /sosmed/{sosmedId} [put]
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

	c.JSON(http.StatusOK, gin.H{
		"message":          "Success Edited Social Media",
		"id":               sosmed.Id,
		"name":             sosmed.Name,
		"social_media_url": sosmed.SocialMediaUrl,
		"user_id":          sosmed.UserId,
	})
}

// GetSosmed Get a social media info
// @Summary Get Sosmed by ID
// @Description Get Sosmed by ID
// @Tags Social Media
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization header in the format of 'Bearer {token}'"
// @Param sosmedId path int true "Sosmed ID"
// @Success 200 {object} helper.SuccessResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 401 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /sosmed/{sosmedId} [get]
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
	// c.JSON(http.StatusOK, sosmed)
	helper.OkWithMessage(c, "Success Get Data", sosmed)
}

// GetSosmed Get all social media info
// @Summary Get all Sosmed
// @Description Get all Sosmed
// @Tags Social Media
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization header in the format of 'Bearer {token}'"
// @Success 200 {object} helper.SuccessResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 401 {object} helper.ErrorResponse
// @Router /sosmed [get]
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
	// c.JSON(http.StatusOK, sosmed)
	helper.OkWithMessage(c, "Success Get All sosmed", sosmed)
}

// DeleteSosmed godoc
// @Summary Delete a social media by ID
// @Description Delete a social media by ID
// @Tags Social Media
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization header in the format of 'Bearer {token}'"
// @Param sosmedId path int true "Sosmed ID"
// @Success 204 "No Content"
// @Failure 400 {object} helper.ErrorResponse
// @Failure 401 {object} helper.ErrorResponse
// @Router /sosmed/{sosmedId} [delete]
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
