package routes

import (
	"context"
	"net/http"

	"flotta-home/mindbond/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type UpdateRequestBody struct {
	UserId   int64  `json:"userId"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Handle   string `json:"handle"`
	Language string `json:"language"`
}

func Update(ctx *gin.Context, c pb.AuthServiceClient) {
	body := UpdateRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Update(context.Background(), &pb.UpdateRequest{
		UserId:   body.UserId,
		Email:    body.Email,
		Password: body.Password,
		Handle:   body.Handle,
		Language: body.Language,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
