package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"

	"go-dart-pub/storage"
)

func New(c *gin.Context) {
	url := location.Get(c)
	c.JSON(http.StatusOK, struct {
		Url    string          `json:"url"`
		Fields json.RawMessage `json:"fields"`
	}{Url: fmt.Sprintf("%v://%v/api/packages/versions/newUpload", url.Scheme, url.Host), Fields: json.RawMessage("{}")})
}

func NewUpload(c *gin.Context) {
	url := location.Get(c)

	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err.Error())
		c.Redirect(302, fmt.Sprintf("%v://%v/api/packages/versions/newUploadFinished?err=%v", url.Scheme, url.Host, err.Error()))
		return
	}

	if err := storage.UploadPackage(file); err != nil {
		log.Println(err.Error())
		c.Redirect(302, fmt.Sprintf("%v://%v/api/packages/versions/newUploadFinished?err=%v", url.Scheme, url.Host, err.Error()))
		return
	}

	c.Redirect(302, fmt.Sprintf("%v://%v/api/packages/versions/newUploadFinished", url.Scheme, url.Host))
}

type successResponse struct {
	Success responseMessage `json:"success"`
}

type errorResponse struct {
	Error responseMessage `json:"error"`
}

type responseMessage struct {
	Message string `json:"message"`
}

func NewUploadFinished(c *gin.Context) {
	err := c.Query("err")
	log.Println(err)
	if err != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{Error: responseMessage{Message: err}})
		return
	}

	c.JSON(http.StatusOK, successResponse{Success: responseMessage{Message: "Successfuly uploaded"}})
}
