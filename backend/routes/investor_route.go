package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vaibhavsijaria/TGC-be.git/handlers"
	"github.com/vaibhavsijaria/TGC-be.git/middleware"
)

func InvesmentRoute(c *gin.Engine) {
	invest := c.Group("/invest", middleware.AuthMiddleware)
	invest.POST("/createinvestment", handlers.OfferInvestment)
	invest.POST("/acceptinvestment", handlers.AccpetToInvestment)
	invest.GET("/investments", handlers.ViewAllInvestment)
}
