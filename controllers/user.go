/*
 * Handle HTTP requests about users
 */

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/noahzyl/gin-ranking/models"
	"github.com/noahzyl/gin-ranking/pkg/logger"
	"strconv"
)

type UserController struct{}

type UserSearch struct {
	Username string `json:"username"`
	Id       int    `json:"id"`
}

func (u *UserController) GetUserInfo(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	user, err := models.GetUser(id)
	if err != nil {
		ReturnErrorJson(ctx, 4001, gin.H{"error": err.Error()})
		return
	}
	ReturnSuccessJson(ctx, 0, "success", user, 1)
}

func (u *UserController) GetUserList(ctx *gin.Context) {
	// Test logger
	logger.Write("Log info...", "user")
	num1 := 5
	num2 := 0
	num3 := num1 / num2
	ReturnErrorJson(ctx, 4001, num3)
}
