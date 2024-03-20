package routes

import (
	"context"
	"flotta-home/mindbond/api-gateway/pkg/chat/pb"
	"net/http"
)

type FindChatRequestBody struct {
	UserId    int64  `json:"userId"`
	ContactId int64  `json:"contactId"`
	Token     string `json:"token"`
}

func FindChat(ctx *gin.Context, c pb.ChatServiceClient) {
	body := FindChatRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.FindChat(context.Background(), &pb.FindChatRequest{
		User1Id: body.UserId,
		User2Id: body.ContactId,
		Token:   body.Token,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
