package auth

import (
	"flotta-home/mindbond/api-gateway/pkg/auth/routes"
	"flotta-home/mindbond/api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routesGroup := r.Group("/auth")
	routesGroup.POST("/register", svc.Register)
	routesGroup.POST("/login", svc.Login)
	routesGroup.POST("/logout", svc.Logout)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}

func (svc *ServiceClient) Logout(ctx *gin.Context) {
	routes.Logout(ctx, svc.Client)
}
