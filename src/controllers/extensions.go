package controllers

import (
	"github.com/gin-gonic/gin"
)

func PubJSON(c *gin.Context, code int, obj interface{}) {
	c.Header("Content-Type", "application/vnd.pub.v2+json")
	c.JSON(code, obj)
}
