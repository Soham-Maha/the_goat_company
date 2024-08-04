package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vaibhavsijaria/TGC-be.git/handlers"
	"github.com/vaibhavsijaria/TGC-be.git/middleware"
)

func ListingRoutes(c *gin.Engine) {
	user := c.Group("/farmer")
	user.POST("/acceptinvestment", middleware.AuthMiddleware, handlers.AcceptInvestment)
	user.POST("/requestinvestment", middleware.AuthMiddleware, handlers.RequestInvestment)
	user.POST("/creategoat", middleware.AuthMiddleware, handlers.CreateGoat)
	user.POST("/sellgoat", middleware.AuthMiddleware, handlers.ListGoatForSale)
	user.POST("/buygoat", middleware.AuthMiddleware, handlers.PurchaseGoat)
	user.GET("/goats", middleware.AuthMiddleware, handlers.ListGoats)
	user.GET("/myorders", middleware.AuthMiddleware, handlers.GetMyOrders)
	user.GET("/mylistings", middleware.AuthMiddleware, handlers.GetMyListings)
	user.GET("/mygoats", middleware.AuthMiddleware, handlers.GetMyGoats)
}
