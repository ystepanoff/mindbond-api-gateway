package routes

import (
	"context"
	"net/http"

	"flotta-home/mindbond/api-gateway/pkg/chat/pb"
)

type CreateChatRequestBody struct {
	User1Id int64  `json:"user1Id"`
	User2Id int64  `json:"user2Id"`
	Token   string `json:"token"`
}

func CreateChat(ctx *gin.Context, c pb.ChatServiceClient) {
	body := CreateChatRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateChat(context.Background(), &pb.CreateChatRequest{
		User1Id: body.User1Id,
		User2Id: body.User2Id,
		Token:   body.Token,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
