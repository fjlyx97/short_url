package main

import "github.com/gin-gonic/gin"

const (
	Address = "0.0.0.0:8080"
)

type SimpleServer struct {
	GinServer *gin.Engine
}

func (s *SimpleServer) Init() {
	s.GinServer = gin.Default()
	s.GinServer.GET("/", s.WebIndex)
	s.GinServer.GET("/:url", s.GetUrl)
	s.GinServer.POST("/", s.SetUrl)
}

func (s *SimpleServer) Run() error {
	return s.GinServer.Run(Address)
}