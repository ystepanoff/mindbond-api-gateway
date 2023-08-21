package routes

import (
	"context"
	"net/http"

	"flotta-home/mindbond/api-gateway/pkg/chat/pb"
	"github.com/gin-gonic/gin"
)

type CreateChatRequestBody struct {
	User1ID int64 `json:"user1Id"`
	User2ID int64 `json:"user2Id"`
}

func CreateChat(ctx *gin.Context, c pb.ChatServiceClient) {
	body := CreateChatRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateChat(context.Background(), &pb.CreateChatRequest{
		User1ID: body.User1ID,
		User2ID: body.User2ID,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
