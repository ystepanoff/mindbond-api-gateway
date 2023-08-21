package routes

import (
	"context"
	"net/http"
	"strconv"

	"flotta-home/mindbond/api-gateway/pkg/chat/pb"
	"github.com/gin-gonic/gin"
)

func FindChat(ctx *gin.Context, c pb.ChatServiceClient) {
	user1Id, _ := strconv.ParseInt(ctx.Param("user1Id"), 10, 32)
	user2Id, _ := strconv.ParseInt(ctx.Param("user2Id"), 10, 32)

	res, err := c.FindChat(context.Background(), &pb.FindChatRequest{
		User1Id: int64(user1Id),
		User2Id: int64(user2Id),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
