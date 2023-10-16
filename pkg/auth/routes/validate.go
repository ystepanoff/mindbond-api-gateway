package routes

import (
	"context"
	"net/http"

	"flotta-home/mindbond/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type ValidateRequestBody struct {
	Token string `json:"token"`
}

func Validate(ctx *gin.Context, c pb.AuthServiceClient) {
	b := ValidateRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Validate(context.Background(), &pb.ValidateRequest{
		Token: b.Token,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
