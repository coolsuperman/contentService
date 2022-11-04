package api

import (
	"contentService/internal/api/content"
	"contentService/pkg/config"
	"contentService/pkg/utils"
	"fmt"
	"github.com/dvwright/xss-mw"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Api struct {
	Gin     *gin.Engine
	config  *config.GConfig
	content *content.ApiContent
}

func NewApi(conf config.GConfig, content *content.ApiContent) *Api {
	return &Api{
		Gin:     gin.Default(),
		config:  &conf,
		content: content,
	}
}

func (a *Api) Run() error {
	r := a.Gin
	var xssMiddleware xss.XssMw
	r.Use(utils.CorsHandler(), xssMiddleware.RemoveXss()) //加载防止跨域中间件,和移除Xss中间件
	contentGroup := r.Group("/content")
	contentGroup.POST("/add_content", a.content.AddContent)
	contentGroup.GET("/delete_content", a.content.DeleteContent)
	contentGroup.GET("/get_content_list", a.content.GetContentList)
	contentGroup.GET("/get_content_detail", a.content.GetContentDetail)
	ginPort := a.config.Api.GinPort
	address := a.config.Api.ListenIP + ":" + strconv.Itoa(ginPort)
	fmt.Println("start api server, address: ", address)
	err := r.Run(address)
	if err != nil {
		fmt.Println("", "run failed ", ginPort, err.Error())
		return err
	}
	return nil
}
