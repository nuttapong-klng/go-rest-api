package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	UserRepo UserRepo
}

type UserRepo interface {
	UserRegistration(user User) error
	UserLogin(loginInfo LoginInfo) (LoginResponse, error)
}

func (h Handlers) UserRegistrationHandlers(context *gin.Context) {
	var requestBody User
	err := context.ShouldBindJSON(&requestBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err = h.UserRepo.UserRegistration(requestBody)
	if err != nil {
		log.Println("UserRepo.UserRegistration", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (h Handlers) UserLoginHandlers(context *gin.Context) {
	var requestBody LoginInfo
	err := context.ShouldBindJSON(&requestBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	loginResponse, err := h.UserRepo.UserLogin(requestBody)
	if err != nil {
		log.Println("UserRepo.UserLogin", err)
		context.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	context.JSON(http.StatusOK, loginResponse)
}
