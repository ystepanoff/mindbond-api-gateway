package routes

import (
	"context"
	"net/http"

	"flotta-home/mindbond/api-gateway/pkg/chat/pb"
	"github.com/gin-gonic/gin"
)

type FetchMessagesRequestBody struct {
	UserId    int64  `json:"userId"`
	ContactId int64  `json:"contactId"`
	Count     int64  `json:"count"`
	Token     string `json:"token"`
}

func FetchMessages(ctx *gin.Context, c pb.ChatServiceClient) {
	body := FetchMessagesRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.FetchMessages(context.Background(), &pb.FetchMessagesRequest{
		UserFromId: body.UserId,
		UserToId:   body.ContactId,
		Count:      body.Count,
		Token:      body.Token,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
