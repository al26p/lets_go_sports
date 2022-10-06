package controllers

import (
	"net/http"

	"github.com/al26p/lets_go_sports/api/models"
	"github.com/gin-gonic/gin"
)

func FindUsers(ctx *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

func CreateUser(ctx *gin.Context) {
	// Validation of the input
	var input models.CreateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create User
	user := models.User{Username: input.Username, LastGame: input.LastGame}
	models.DB.Create(&user)

	//Todo : Catch DB errrors

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func FindUser(ctx *gin.Context) {
	var user models.User

	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func PatchUser(ctx *gin.Context) {
	var user models.User

	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.CreateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create User
	models.DB.Model(&user).Updates(input)

	//Todo : Catch DB errrors
	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(ctx *gin.Context) {
	var user models.User

	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&user)

	ctx.JSON(http.StatusOK, gin.H{"data": true})
}
