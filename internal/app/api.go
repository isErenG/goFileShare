package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/upload", uploadFile)

	r.GET("/download", downloadFile)

	return r
}

func uploadFile(c *gin.Context) {
	fmt.Println("Test")
}

func downloadFile(c *gin.Context) {
	fmt.Println("Test")
}
