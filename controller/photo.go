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

// CreatePhoto Create a new Photo
// @Summary Create a new Photo
// @Description Create a new Photo
// @Tags Photo
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization header in the format of 'Bearer {token}'"
// @Param photo body model.PhotoReqEq true "Sosmed data"
// @Success 201 {object} helper.SuccessResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /photos [post]
func CreatePhoto(c *gin.Context) {
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

	// c.JSON(http.StatusCreated, photo)
	helper.DataCreated(c, "Photo Successfully Created", photo)
}

// UpdatePhoto Update a photo info
// @Summary Update existing photo
// @Description Update existing photo
// @Tags Photo
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization header in the format of 'Bearer {token}'"
// @Param photoId path int true "Photo ID"
// @Param photo body model.UpdatePhotoReq true "Update photo object"
// @Success 200 {object} helper.SuccessResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 401 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /photo/{photoId} [put]
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

	// c.JSON(http.StatusOK, photo)
	helper.OkWithMessage(c, "Photo Updated", photo)
}

// GetPhoto Get a photo info
// @Summary Get photo by ID
// @Description Get photo by ID
// @Tags Photo
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization header in the format of 'Bearer {token}'"
// @Param photoId path int true "photo ID"
// @Success 200 {object} helper.SuccessResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 401 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /photo/{photoId} [get]
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
	// c.JSON(http.StatusOK, photo)
	helper.OkWithMessage(c, "Succes Get Photo", photo)
}

// GetPhoto Get all photo info
// @Summary Get all photo
// @Description Get all photo
// @Tags Photo
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization header in the format of 'Bearer {token}'"
// @Success 200 {object} helper.SuccessResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 401 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /photo [get]
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
	// c.JSON(http.StatusOK, photo)
	helper.OkWithMessage(c, "Success Get All photo", photo)
}

// DeletePhoto godoc
// @Summary Delete a photo by ID
// @Description Delete a photo by ID
// @Tags Photo
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization header in the format of 'Bearer {token}'"
// @Param photoId path int true "Photo ID"
// @Success 204 "No Content"
// @Failure 400 {object} helper.ErrorResponse
// @Failure 401 {object} helper.ErrorResponse
// @Router /photo/{photoId} [delete]
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
