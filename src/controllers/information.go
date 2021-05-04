package controllers

import (
	"log"
	"net/http"

	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"

	"go-dart-pub/storage"
)

func FindPackageVersions(c *gin.Context) {

	packageName := c.Param("packageName")
	pi, err := storage.FindPackageVersions(packageName, location.Get(c))

	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "package named " + packageName + " not found",
		})
		return
	}

	PubJSON(c, http.StatusOK, pi)
}
