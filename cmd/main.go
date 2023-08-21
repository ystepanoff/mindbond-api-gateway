package main

import (
	"log"

	"flotta-home/mindbond/api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	// authSvc := *auth.RegisterRoutes(r, &c)
	// chat.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
