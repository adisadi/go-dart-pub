package controllers

import (
	"fmt"
	"go-dart-pub/storage"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func DownloadPackage(c *gin.Context) {

	packageName := c.Param("packageName")
	version := c.Param("version")

	version = strings.Replace(version, ".tar.gz", "", -1)

	filePath, err := storage.Download(packageName, version)

	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "package named " + packageName + " not found",
		})
		return
	}

	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s-%s.tar.gz", packageName, version))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(filePath)
}
