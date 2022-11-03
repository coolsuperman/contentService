package main

import (
	"contentService/internal/api/content"
	"contentService/pkg/config"
	"contentService/pkg/utils"
	"fmt"
	"github.com/dvwright/xss-mw"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	r := gin.Default()
	var xssMiddleware xss.XssMw
	r.Use(utils.CorsHandler(), xssMiddleware.RemoveXss()) //加载防止跨域中间件,和移除Xss中间件
	contentGroup := r.Group("/content")
	contentGroup.POST("/add_content", content.AddContent)
	contentGroup.GET("/delete_content", content.DeleteContent)
	contentGroup.GET("/get_content_list", content.GetContentList)
	contentGroup.GET("/get_content_detail", content.GetContentDetail)
	ginPort := config.Config.Api.GinPort
	address := config.Config.Api.ListenIP + ":" + strconv.Itoa(ginPort)
	fmt.Println("start api server, address: ", address)
	err := r.Run(address)
	if err != nil {
		fmt.Println("", "run failed ", ginPort, err.Error())
	}
}
