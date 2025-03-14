/*
 * Handle HTTP requests about users
 */

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/noahzyl/gin-ranking/models"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

type UserController struct{}

// UserRequest struct is used to save parameters from the client which are sent in JSON
type UserRequest struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// User struct sent to the client
type UserResponse struct {
	Id       int
	Username string
}

func (u *UserController) AddUser(ctx *gin.Context) {
	user := &UserRequest{}
	err := ctx.BindJSON(user) // Parse parameters in the request
	if err != nil {
		ReturnErrorJson(ctx, 4001, gin.H{"Parameters parsing error": err.Error()})
		return
	}
	// Use bcrypt to hash user's password
	var dbErr error
	hashedPassword, cryptErr := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if cryptErr != nil {
		ReturnErrorJson(ctx, 4002, gin.H{"Adding user failed": cryptErr.Error()})
		return
	}
	// Add a new user to the database
	var id int
	id, dbErr = models.AddUser(user.Username, string(hashedPassword))
	if dbErr != nil {
		ReturnErrorJson(ctx, 4002, gin.H{"Adding user failed": dbErr.Error()})
		return
	}
	ReturnSuccessJson(ctx, 0, "Adding user succeed", id, 1)
}

func (u *UserController) UpdateUserName(ctx *gin.Context) {
	user := &UserRequest{}
	err := ctx.BindJSON(user)
	if err != nil {
		ReturnErrorJson(ctx, 4001, gin.H{"Parameters parsing error": err.Error()})
		return
	}
	var dbErr error
	// Update username
	if user.Username != "" {
		dbErr = models.UpdateUserName(user.Id, user.Username)
		if dbErr != nil {
			ReturnErrorJson(ctx, 4002, gin.H{"Updating username failed": dbErr.Error()})
			return
		}
		ReturnSuccessJson(ctx, 0, "Updating username succeed", true, 1)
	}
	// Update password
	if user.Password != "" {
		hashedPassword, cryptErr := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if cryptErr != nil {
			ReturnErrorJson(ctx, 4002, gin.H{"Updating password failed": cryptErr.Error()})
			return
		}
		dbErr = models.UpdateUserPassword(user.Id, string(hashedPassword))
		if dbErr != nil {
			ReturnErrorJson(ctx, 4002, gin.H{"Updating password failed": dbErr.Error()})
			return
		}
		ReturnSuccessJson(ctx, 0, "Updating password succeed", true, 1)
	}
}

func (u *UserController) DeleteUser(ctx *gin.Context) {
	user := &UserRequest{}
	err := ctx.BindJSON(user)
	if err != nil {
		ReturnErrorJson(ctx, 4001, gin.H{"Parameters parsing error": err.Error()})
		return
	}
	dbErr := models.DeleteUser(user.Id)
	if dbErr != nil {
		ReturnErrorJson(ctx, 4002, gin.H{"Deleting user failed": dbErr.Error()})
		return
	}
	ReturnSuccessJson(ctx, 0, "Deleting user succeed", true, 1)
}

func (u *UserController) GetUserInfo(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	user, dbErr := models.GetUser(id)
	if dbErr != nil {
		ReturnErrorJson(ctx, 4004, gin.H{"No user found": dbErr.Error()})
		return
	}
	userResponse := UserResponse{Id: user.Id, Username: user.Username}
	ReturnSuccessJson(ctx, 0, "Getting user succeed", userResponse, 1)
}

func (u *UserController) GetUserList(ctx *gin.Context) {
	username := ctx.Param("username")
	users, dbErr := models.GetUserByUsername(username)
	if dbErr != nil {
		ReturnErrorJson(ctx, 4002, gin.H{"Getting users failed": dbErr.Error()})
		return
	}
	var usersResponseList []UserResponse
	for _, user := range users {
		userResponse := UserResponse{Id: user.Id, Username: user.Username}
		usersResponseList = append(usersResponseList, userResponse)
	}
	ReturnSuccessJson(ctx, 0, "Getting users succeed", usersResponseList, 1)
}
