package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler interface {
	AddUser(ctx context.Context, req *AddUserReq) (*AddUserReply, error)
}

func HttpHandler(g *gin.Engine, us UserHandler) http.Handler {
	g.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "ok",
		})
	})

	g.POST("user.v1/AddUser", func(context *gin.Context) {
		req := &AddUserReq{
			name: context.PostForm("name"),
		}
		rep, err := us.AddUser(context, req)
		if err != nil {
			context.JSON(500, err.Error())
		}
		context.JSON(200, rep)
	})

	return g
}
