package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *SimpleServer) WebIndex(c *gin.Context) {
	c.String(200, "Okk")
}

func (s *SimpleServer) SetUrl(c *gin.Context) {
	url := c.PostForm("url")
	// 参数校验
	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Url is empty",
		})
		return
	}
	shortUrl, err := s.SetUrlService(c, url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Bad service",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"url": shortUrl,
	})
}

func (s *SimpleServer) GetUrl(c *gin.Context) {
	url := c.Param("url")
	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Url is empty",
		})
		return
	}
	longUrl, err := s.GetUrlService(c, url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Bad service",
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, longUrl)
}
