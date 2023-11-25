package routes

import (
	"context"
	"net/http"

	"flotta-home/mindbond/api-gateway/pkg/chat/pb"
	"github.com/gin-gonic/gin"
)

type AddMessageRequestBody struct {
	UserFromId       int64  `json:"userFromId"`
	UserFromLanguage string `json:"userFromLanguage"`
	UserToId         int64  `json:"userToId"`
	UserToLanguage   string `json:"userToLanguage"`
	Message          string `json:"message"`
}

func AddMessage(ctx *gin.Context, c pb.ChatServiceClient) {
	body := AddMessageRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.AddMessage(context.Background(), &pb.AddMessageRequest{
		UserFromId:       body.UserFromId,
		UserFromLanguage: body.UserFromLanguage,
		UserToId:         body.UserToId,
		UserToLanguage:   body.UserToLanguage,
		Message:          body.Message,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
