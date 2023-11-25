package routes

import (
	"context"
	"net/http"

	"flotta-home/mindbond/api-gateway/pkg/chat/pb"
	"github.com/gin-gonic/gin"
)

type AddContactRequestBody struct {
	UserId int64  `json:"userId"`
	Token  string `json:"token"`
	Handle string `json:"handle"`
}

func AddContact(ctx *gin.Context, c pb.ChatServiceClient) {
	body := AddContactRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.AddContact(context.Background(), &pb.AddContactRequest{
		UserId: body.UserId,
		Token:  body.Token,
		Handle: body.Handle,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
