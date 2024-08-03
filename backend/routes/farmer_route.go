package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vaibhavsijaria/TGC-be.git/handlers"
	"github.com/vaibhavsijaria/TGC-be.git/middleware"
)

func ListingRoutes(c *gin.Engine) {
	user := c.Group("/farmer")
	user.POST("/creategoat", middleware.AuthMiddleware, handlers.CreateGoat)
	user.GET("/goats", middleware.AuthMiddleware, handlers.ListGoats)

}
