package routes

import (
	"log"
	"net/http"

	"example.com/rest-api/modells"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user modells.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't parse data"})
		return
	}
	err := user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't save user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})

}

func getUsers(context *gin.Context) {
	users, err := modells.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't fetch events"})
		return
	}
	context.JSON(http.StatusOK, users)
}

func login(context *gin.Context) {
	var user modells.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid input"})
		return
	}

	log.Println("User trying to login with email:", user.Email)
	log.Println("User trying to login with password:", user.Password)

	check, err := user.ValidateUser()
	log.Println("User validation check:", check)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error while validating user"})
		return
	}
	if !check {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	log.Println("User validated with ID:", user.ID)

	token, err := utils.GenerateJWTToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't generate token"})
		return
	}

	context.JSON(200, gin.H{"message": "Login successful", "user_id": user.ID, "token": token})
}
