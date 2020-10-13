package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	Address = "0.0.0.0:8080"
)

type SimpleServer struct {
	GinServer *gin.Engine
}

func (s *SimpleServer) Init() {
	s.GinServer = gin.Default()
	s.GinServer.LoadHTMLFiles("./views/index.html")
	s.GinServer.StaticFS("/static/", http.Dir("./views"))
	s.GinServer.StaticFile("/favicon.ico", "./views/favicon.ico")
	s.GinServer.GET("/", s.WebIndex)
	//s.GinServer.GET("/:url", s.GetUrl)
	s.GinServer.NoRoute(s.NoRoute)
	s.GinServer.POST("/", s.SetUrl)
}

func (s *SimpleServer) Run() error {
	return s.GinServer.Run(Address)
}
