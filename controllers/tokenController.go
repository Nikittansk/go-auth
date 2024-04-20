package controllers

import (
	"net/http"

	"github.com/Nikittansk/go-auth/auth"
	"github.com/Nikittansk/go-auth/database"
	"github.com/Nikittansk/go-auth/models"
	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateToken(ctx *gin.Context) {
	var request TokenRequest
	var user models.User

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	if err := database.DB.QueryRow("SELECT * FROM users WHERE email=$1", request.Email).Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	if err := user.CheckPassword(request.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	tokenString, err := auth.GenerateJWT(user.Username, user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
