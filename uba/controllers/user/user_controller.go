package user

import (
	"net/http"
	"strconv"
	dto2 "uba/dto"
	service "uba/services"

	_ "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetUserByEmail(c *gin.Context) {
	var userDto dto2.UserDto
	c.JSON(http.StatusOK, userDto)

}

func GetUserById(c *gin.Context) {
	log.Debug("User ID to upload: " + c.Param("id"))
	id, _ := strconv.Atoi(c.Param("id"))
	var userDto dto2.UserDto

	userDto, err := service.UserService.GetUserById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, userDto)
}

//login

func UserLogin(c *gin.Context) {
	var loginDto dto2.LoginDto
	err := c.BindJSON(&loginDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	tokenDto, er := service.UserService.LoginUser(loginDto)

	if er != nil {
		c.JSON(er.Status(), er)
		return
	}
	c.JSON(http.StatusOK, tokenDto)

}

func InsertUser(c *gin.Context) {
	var userDto dto2.UserDto
	err := c.BindJSON(&userDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	tokenDto, er := service.UserService.InsertUser(userDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, tokenDto)
}
