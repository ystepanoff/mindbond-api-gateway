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
	routesGroup.POST("/add_contact", svc.AddContact)
	routesGroup.POST("/remove_contact", svc.RemoveContact)
	routesGroup.POST("/fetch_contacts", svc.FetchContacts)
	routesGroup.POST("/create", svc.CreateChat)
	routesGroup.POST("/send", svc.AddMessage)
	routesGroup.GET("/find/:user1Id/:user2Id", svc.FindChat)
}

func (svc *ServiceClient) AddContact(ctx *gin.Context) {
	routes.AddContact(ctx, svc.Client)
}

func (svc *ServiceClient) RemoveContact(ctx *gin.Context) {
	routes.RemoveContact(ctx, svc.Client)
}

func (svc *ServiceClient) FetchContacts(ctx *gin.Context) {
	routes.FetchContacts(ctx, svc.Client)
}

func (svc *ServiceClient) CreateChat(ctx *gin.Context) {
	routes.CreateChat(ctx, svc.Client)
}

func (svc *ServiceClient) FindChat(ctx *gin.Context) {
	routes.FindChat(ctx, svc.Client)
}

func (svc *ServiceClient) AddMessage(ctx *gin.Context) {
	routes.AddMessage(ctx, svc.Client)
}
