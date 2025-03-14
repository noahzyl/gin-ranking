/*
 * Define routers
 */

package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/noahzyl/gin-ranking/controllers"
	"github.com/noahzyl/gin-ranking/pkg/logger"
)

func Router() *gin.Engine {
	r := gin.Default() // Create a router

	// Set a logger
	r.Use(gin.LoggerWithConfig(logger.LogRequest()))
	r.Use(logger.LogError)

	// Set a router group of users (user router)
	user := r.Group("/user")
	{
		// controllers.UserController{} will create an anonymous variable of UserController
		user.POST("/add", (&controllers.UserController{}).AddUser) // // ctx is automatically passed by gin
		user.POST("/update", (&controllers.UserController{}).UpdateUserName)
		user.POST("/delete", (&controllers.UserController{}).DeleteUser)
		user.GET("/info/:id", (&controllers.UserController{}).GetUserInfo)
		user.GET("/list/:username", (&controllers.UserController{}).GetUserList)
	}

	// Set a router group of ranking (ranking router)
	ranking := r.Group("/ranking")
	{
		ranking.POST("/list", (&controllers.RankingController{}).GetRankingList)
	}

	return r
}
