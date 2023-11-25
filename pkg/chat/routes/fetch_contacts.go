package routes

import (
	"context"
	"net/http"

	"flotta-home/mindbond/api-gateway/pkg/chat/pb"
	"github.com/gin-gonic/gin"
)

type FetchContactsRequestBody struct {
	UserId int64  `json:"userId"`
	Token  string `json:"token"`
}

func FetchContacts(ctx *gin.Context, c pb.ChatServiceClient) {
	body := FetchContactsRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.FetchContacts(context.Background(), &pb.FetchContactsRequest{
		UserId: body.UserId,
		Token:  body.Token,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
