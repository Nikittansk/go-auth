package controllers

import (
	"net/http"

	"github.com/Nikittansk/go-auth/database"
	"github.com/Nikittansk/go-auth/models"
	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	if err := user.HashPassword(user.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	_, err := database.DB.Exec("INSERT INTO users (name, username, email, password) VALUES ($1, $2, $3, $4)", user.Name, user.Username, user.Email, user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, user)
}
