package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type Server_2 struct {
	httpServer *http.Server
}

func NewService2() *Server_2 {
	g := gin.Default()
	g.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	return &Server_2{
		httpServer: &http.Server{
			Addr:    ":8081",
			Handler: g,
		},
	}
}

func (s *Server_2) Start() error {
	log.Println("service2 启动")
	return s.httpServer.ListenAndServe()
}

func (s *Server_2) Stop() error {
	log.Println("service2 关闭")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return s.httpServer.Shutdown(ctx)
}
