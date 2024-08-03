package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vaibhavsijaria/TGC-be.git/handlers"
	"github.com/vaibhavsijaria/TGC-be.git/middleware"
)

func UserRoutes(c *gin.Engine) {
	user := c.Group("/user")
	user.POST("/signup", handlers.Signup)
	user.POST("/login", handlers.Login)
	user.GET("/logincheck", middleware.AuthMiddleware, handlers.LoginCheck)
}
