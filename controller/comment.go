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

func CreateComment(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)

	comment := model.Comment{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&comment)
	} else {
		c.ShouldBind(&comment)
	}

	comment.UserId = userID

	err := db.Debug().Create(&comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, comment)
}

func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)

	commentId, _ := strconv.Atoi(c.Param("commentId"))

	comment := model.Comment{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&comment)
	} else {
		c.ShouldBind(&comment)
	}

	comment.UserId = userID
	comment.Id = uint(commentId)

	err := db.Model(&comment).Where("id = ?", commentId).Updates(model.Comment{
		PhotoId: comment.PhotoId,
		Message: comment.Message,
	}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, comment)
}

func GetComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	commentId, _ := strconv.Atoi(c.Param("commentId"))

	comment := model.Comment{}
	userID := uint(userData["id"].(float64))

	comment.UserId = userID
	comment.Id = uint(commentId)

	err := db.Preload("User", func(db_s *gorm.DB) *gorm.DB {
		return db_s.Select("id, user_name, age")
	}).First(&comment, commentId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})

		return
	}
	c.JSON(http.StatusOK, comment)
}

func GetAllCcomment(c *gin.Context) {
	db := database.GetDB()
	comment := []model.Comment{}

	err := db.Preload("User", func(db_s *gorm.DB) *gorm.DB {
		return db_s.Select("id, user_name, age")
	}).Find(&comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})

		return
	}
	c.JSON(http.StatusOK, comment)
}

func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	comment := model.Comment{}

	comment.UserId = userID
	comment.Id = uint(commentId)

	err := db.Delete(&comment, commentId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})

		return
	}
	c.JSON(http.StatusNoContent, nil)
}
