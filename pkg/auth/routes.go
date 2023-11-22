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
	routesGroup.POST("/update", svc.Update)
	routesGroup.POST("/login", svc.Login)
	routesGroup.POST("/dry-login", svc.DryLogin)
	routesGroup.POST("/logout", svc.Logout)
	routesGroup.POST("/validate", svc.Validate)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Update(ctx *gin.Context) {
	routes.Update(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}

func (svc *ServiceClient) DryLogin(ctx *gin.Context) {
	routes.DryLogin(ctx, svc.Client)
}

func (svc *ServiceClient) Logout(ctx *gin.Context) {
	routes.Logout(ctx, svc.Client)
}

func (svc *ServiceClient) Validate(ctx *gin.Context) {
	routes.Validate(ctx, svc.Client)
}
