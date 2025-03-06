/*
 * Handle HTTP requests about users
 */

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/noahzyl/gin-ranking/pkg/logger"
)

type UserController struct{}

type UserSearch struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

func (u *UserController) GetUserInfo(ctx *gin.Context) {
	ReturnSuccessJson(ctx, 0, "Success", "/user/info...", 1)
}

func (u *UserController) GetUserList(ctx *gin.Context) {
	//search := &UserSearch{}
	//err := ctx.BindJSON(search) // Parse parameters of the request
	//if err != nil {
	//	ReturnErrorJson(ctx, 4001, gin.H{"error": err})
	//}
	//ReturnSuccessJson(ctx, 0, search.Name, search.ID, 1)

	// Realize a panic and recover it
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println("error:", err)
	//	}
	//}()

	// Test logger
	logger.Write("Log info...", "user")
	num1 := 5
	num2 := 0
	num3 := num1 / num2
	ReturnErrorJson(ctx, 4001, num3)
}
