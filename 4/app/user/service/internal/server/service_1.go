package server

import (
	"context"
	"github.com/gin-gonic/gin"
	v1 "go_crouse/4/api/user/v1"
	"go_crouse/4/app/user/service/internal/service"
	"log"
	"net/http"
	"time"
)

type s1 struct {
	httpServer *http.Server
}

func NewService1(us *service.UserService) *s1 {
	g := gin.Default()
	handler := v1.HttpHandler(g, us)

	return &s1{
		httpServer: &http.Server{
			Addr:    ":8080",
			Handler: handler,
		},
	}
}

func (s *s1) Start() error {
	log.Println("service1 启动")
	return s.httpServer.ListenAndServe()
}

func (s *s1) Stop() error {
	log.Println("service1 关闭")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return s.httpServer.Shutdown(ctx)
}
