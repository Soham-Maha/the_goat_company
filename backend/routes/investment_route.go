package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vaibhavsijaria/TGC-be.git/handlers"
	"github.com/vaibhavsijaria/TGC-be.git/middleware"
)

func InvesmentRoute(c *gin.Engine) {
	user := c.Group("/invest")
	user.POST("/createinvestment", middleware.AuthMiddleware, handlers.CreateInvestment)
	// user.GET("/investments", middleware.AuthMiddleware, handlers.ListGoats)

}
