package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init()  {
	router = gin.Default()
	router.Use(cors.Default())
	//router.Run()
}

func StartApp()  {
	mapUrls()
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
