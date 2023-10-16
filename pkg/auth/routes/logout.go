package routes

import (
	"context"
	"net/http"

	"flotta-home/mindbond/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type LogoutRequestBody struct {
	UserId int64  `json:"id"`
	Token  string `json:"token"`
}

func Logout(ctx *gin.Context, c pb.AuthServiceClient) {
	b := LogoutRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Logout(context.Background(), &pb.LogoutRequest{
		UserId: b.UserId,
		Token:  b.Token,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
