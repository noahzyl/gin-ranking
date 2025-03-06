/*
 * Define routers
 */

package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/noahzyl/gin-ranking/controllers"
	"github.com/noahzyl/gin-ranking/pkg/logger"
	"net/http"
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
		user.GET("/info", (&controllers.UserController{}).GetUserInfo) // ctx is automatically passed by gin
		user.POST("/list", (&controllers.UserController{}).GetUserList)
		user.PUT("/add", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "/user/add...")
		})
		user.DELETE("/delete", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "/user/delete...")
		})
	}

	// Set a router group of ranking (ranking router)
	ranking := r.Group("/ranking")
	{
		ranking.POST("/list", (&controllers.RankingController{}).GetRankingList)
	}

	return r
}
