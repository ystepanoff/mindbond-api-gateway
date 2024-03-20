package routes

import (
	"context"
	"net/http"

	"flotta-home/mindbond/api-gateway/pkg/chat/pb"
)

type AddMessageRequestBody struct {
	UserFromId int64  `json:"userFromId"`
	UserToId   int64  `json:"userToId"`
	Message    string `json:"message"`
	Token      string `json:"token"`
}

func AddMessage(ctx *gin.Context, c pb.ChatServiceClient) {
	body := AddMessageRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.AddMessage(context.Background(), &pb.AddMessageRequest{
		UserFromId: body.UserFromId,
		UserToId:   body.UserToId,
		Message:    body.Message,
		Token:      body.Token,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
