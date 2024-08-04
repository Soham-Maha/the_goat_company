package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vaibhavsijaria/TGC-be.git/handlers"
	"github.com/vaibhavsijaria/TGC-be.git/middleware"
)

func ListingRoutes(c *gin.Engine) {
	farmer := c.Group("/farmer", middleware.AuthMiddleware)

	farmer.POST("/acceptinvestment", handlers.AcceptInvestment)
	farmer.POST("/requestinvestment", handlers.RequestInvestment)
	farmer.POST("/creategoat", handlers.CreateGoat)
	farmer.POST("/sellgoat", handlers.ListGoatForSale)
	farmer.POST("/buygoat", handlers.PurchaseGoat)

	farmer.GET("/goats", handlers.ListGoats)
	farmer.GET("/myorders", handlers.GetMyOrders)
	farmer.GET("/mylistings", handlers.GetMyListings)
	farmer.GET("/mygoats", handlers.GetMyGoats)
	farmer.GET("/myinvestmentrequest", handlers.FarmerInvestmentReq)
}
