package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"

	"go-dart-pub/controllers"
	"go-dart-pub/storage"
)

func main() {

	setStorageBasePath()

	r := gin.Default()

	r.Use(location.New(location.Config{
		Headers: location.Headers{Scheme: "X-Forwarded-Proto", Host: "X-Forwarded-Host"},
		Scheme:"http"
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"info": "Welcome to go-dart-pub, a dart pub microservice"})
	})

	//Package Information
	r.GET("/api/packages/:packageName", controllers.FindPackageVersions)

	//Package Download
	r.GET("/packages/:packageName/versions/:version", controllers.DownloadPackage)

	//Package Upload
	r.GET("/api/packages/versions/new", controllers.New)
	r.POST("/api/packages/versions/newUpload", controllers.NewUpload)
	r.GET("/api/packages/versions/newUploadFinished", controllers.NewUploadFinished)

	r.Run()
}

func setStorageBasePath() {
	base := os.Getenv("DART_PUB_STORAGE_BASE")
	if base != "" {
		storage.StorageBasePath = base
	}
	log.Printf("Base Path= '%s'\n", storage.StorageBasePath)
}
