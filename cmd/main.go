package main

import (
	"log"

	"flotta-home/mindbond/api-gateway/pkg/auth"
	"flotta-home/mindbond/api-gateway/pkg/chat"
	"flotta-home/mindbond/api-gateway/pkg/config"

	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	router := gin.Default()

	// CORS
	router.Use(func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, Authorization")
		ctx.Header("Access-Control-Allow-Methods", "POST, OPTIONS")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()
	})

	authService := *auth.RegisterRoutes(router, &c)
	chat.RegisterRoutes(router, &c, &authService)

	router.Run(c.Port)
}
