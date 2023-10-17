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

	r := gin.Default()

	authService := *auth.RegisterRoutes(r, &c)
	chat.RegisterRoutes(r, &c, &authService)

	r.Run(c.Port)
}
