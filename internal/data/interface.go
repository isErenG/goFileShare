package data

import "github.com/gin-gonic/gin"

type Services interface {
	uploadFile(c *gin.Context)
	downloadFile(c *gin.Context)
}
