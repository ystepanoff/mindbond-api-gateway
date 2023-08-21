package chat

import (
	"flotta-home/mindbond/api-gateway/pkg/auth"
	"flotta-home/mindbond/api-gateway/pkg/chat/routes"
	"flotta-home/mindbond/api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routesGroup := r.Group("/chat")
	routesGroup.Use(a.AuthRequired)
	routesGroup.POST("/", svc.CreateChat)
	routesGroup.GET("/:user1Id/:user2Id", svc.FindChat)
}

func (svc *ServiceClient) CreateChat(ctx *gin.Context) {
	routes.CreateChat(ctx, svc.Client)
}

func (svc *ServiceClient) FindChat(ctx *gin.Context) {
	routes.FindChat(ctx, svc.Client)
}
