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

// CreateComment Create a new Comment
// @Summary Create a new Comment
// @Description Create a new Comment
// @Tags Comment
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization header in the format of 'Bearer {token}'"
// @Param comment body model.CommentReqEq true "Comment data"
// @Success 201 {object} helper.SuccessResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /comment [post]
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

	// c.JSON(http.StatusCreated, comment)
	helper.DataCreated(c, "Comment Successfully Created", comment)
}

// UpdateComment Update a comment info
// @Summary Update existing comment
// @Description Update existing comment
// @Tags Comment
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization header in the format of 'Bearer {token}'"
// @Param commentId path int true "comment ID"
// @Param comment body model.UpdateCommentReq true "Update Comment object"
// @Success 200 {object} helper.SuccessResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 401 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /comment/{commentId} [put]
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

	// c.JSON(http.StatusOK, comment)
	helper.OkWithMessage(c, "Succes Update Comment", comment)
}

// GetComment Get a comment info
// @Summary Get comment by ID
// @Description Get comment by ID
// @Tags Comment
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization header in the format of 'Bearer {token}'"
// @Param commentId path int true "Comment ID"
// @Success 200 {object} helper.SuccessResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 401 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /comment/{commentId} [get]
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
	// c.JSON(http.StatusOK, comment)
	helper.OkWithMessage(c, "Succes Get Comment", comment)
}

// GetComment Get all comment info
// @Summary Get all comment
// @Description Get all comment
// @Tags Comment
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization header in the format of 'Bearer {token}'"
// @Success 200 {object} helper.SuccessResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 401 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /comment [get]
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
	// c.JSON(http.StatusOK, comment)
	helper.OkWithMessage(c, "Success Get All Comment", comment)
}

// DeleteComment godoc
// @Summary Delete a comment by ID
// @Description Delete a comment by ID
// @Tags Comment
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization header in the format of 'Bearer {token}'"
// @Param commentId path int true "Comment ID"
// @Success 204 "No Content"
// @Failure 400 {object} helper.ErrorResponse
// @Failure 401 {object} helper.ErrorResponse
// @Router /comment/{commentId} [delete]
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
