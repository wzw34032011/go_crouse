package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type s1 struct {
	httpServer *http.Server
}

func NewService1() *s1 {
	//todo 通过service加载handler
	g := gin.Default()
	g.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	return &s1{
		httpServer: &http.Server{
			Addr:    ":8080",
			Handler: g,
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
