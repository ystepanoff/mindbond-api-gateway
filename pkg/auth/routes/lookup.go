package routes

import (
	"context"
	"net/http"

	"flotta-home/mindbond/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type LookupRequestBody struct {
	UserId int64  `json:"userId"`
	Email  string `json:"email"`
	Handle string `json:"handle"`
}

func Lookup(ctx *gin.Context, c pb.AuthServiceClient) {
	b := LookupRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Lookup(context.Background(), &pb.LookupRequest{
		UserId: b.UserId,
		Email:  b.Email,
		Handle: b.Handle,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
