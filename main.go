package main

import (
	"github.com/gin-gonic/gin"
	"auth-system-go/handlers"
	"auth-system-go/middleware"
)

func main(){
	r :=gin.Default()

	// check health
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	r.POST("register", handlers.Register)
	r.POST("login", handlers.Login)

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("profile", handlers.GetProfile)
	}

	// run server
	r.Run(":8080")
}
