package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vaibhavsijaria/TGC-be.git/handlers"
	"github.com/vaibhavsijaria/TGC-be.git/middleware"
)

func VetRoute(c *gin.Engine) {
	user := c.Group("/vet")
	user.POST("/healthcheck", middleware.AuthMiddleware, handlers.HealthCheck)
}
