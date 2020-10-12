package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (s *SimpleServer) WebIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
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

func (s *SimpleServer) GetUrl(c *gin.Context, url string) {
	//url := c.Param("url")
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

// 自定义路由，gin自带路由无法匹配特殊情况
func (s *SimpleServer) NoRoute(c *gin.Context) {
	url := c.Request.URL.String()
	// 判断\字符出现次数
	charNums := strings.Count(url, "/")
	fmt.Println(charNums)
	if len(url) > 20 || charNums != 1 {
		c.Redirect(http.StatusMovedPermanently, "/")
	}
	s.GetUrl(c, url)
}
