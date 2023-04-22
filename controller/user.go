package controller

import (
	"my-gram/database"
	"my-gram/helper"
	"my-gram/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

// Register a new user
// @Summary Register a new user
// @Description Register a new user
// @Tags Action For Users
// @Accept json
// @Produce json
// @Param user body model.UserReqEq true "User object that needs to be registered"
// @Success 201 {object} helper.UserRegisterResponse
// @Failure 400 {object} helper.ErrorResponse
// @Router /users/register [post]
func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	user := model.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	err := db.Debug().Create(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":   "Success Created User",
		"id":        user.Id,
		"email":     user.Email,
		"full_name": user.UserName,
	})
}

// UserLogin godoc
// @Summary User login
// @Description User login API
// @Tags Action For Users
// @Accept  json
// @Produce  json
// @Param user body model.LoginReq true "User login credentials"
// @Success 200 {object} helper.UserLoginResponse
// @Failure 401 {object} helper.ErrorResponse
// @Router /users/login [post]
func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	user := model.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	passwordClient := user.Password

	err := db.Debug().Where("email = ?", user.Email).Take(&user).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err":     "unauthorized",
			"message": "invalid email or password",
		})
		return
	}

	isValid := helper.ComparePass([]byte(user.Password), []byte(passwordClient))
	if !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err":     "unauthorized",
			"message": "invalid email or password",
		})
		return
	}

	token := helper.GenerateToken(user.Id, user.Email)

	c.JSON(http.StatusOK, gin.H{
		"message": "User successfully logged in",
		"token":   token,
	})
}
