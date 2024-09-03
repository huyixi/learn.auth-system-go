package main

import (
	"github.com/gin-gonic/gin"
)

func main(){
	r :=gin.Default()

	// check health
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	r.POST("register", Register)
	r.POST("login", Login)
	r.GET("profile", AuthMiddleware(), Profile)

	// run server
	r.Run()
}
