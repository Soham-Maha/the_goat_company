package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vaibhavsijaria/TGC-be.git/handlers"
	"github.com/vaibhavsijaria/TGC-be.git/middleware"
)

func ListingRoutes(c *gin.Engine) {
	user := c.Group("/create")
	user.POST("/goat", middleware.AuthMiddleware, handlers.CreateGoat)
}
