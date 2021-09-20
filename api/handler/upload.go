package handler

import (
	"fmt"
	"net/http"

	"gogogo/pkg/errors"
	"gogogo/pkg/response"
	"gogogo/pkg/success"

	"gogogo/entity/responses"
	"gogogo/repository"
	"io"
	"strings"

	"log"

	"github.com/gin-gonic/gin"
)

func HandleFileupload(c *gin.Context) {
	file, handler, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": errors.ErrParsingFile.Error(),
			"error":   true,
		})
		return
	}
	defer file.Close()
	filename := handler.Filename
	const MaxFileSize = 8 * 1000000 //8 megabytes (in bytes)
	filetype := handler.Header.Get("Content-Type")
	if !strings.Contains(filetype, "image") {
		c.JSON(http.StatusForbidden, gin.H{
			"message": errors.ErrFalseContent.Error(),
			"error":   true,
		})
		return
	}
	fmt.Println("contentsize", handler.Size)
	fmt.Println("contentsize", MaxFileSize)
	if handler.Size > MaxFileSize {
		c.JSON(http.StatusRequestEntityTooLarge, gin.H{
			"message": errors.ErrTooLarge.Error(),
			"error":   true,
		})
		return
	}
	firebaseStorage := repository.FirebaseStorage()
	bucket, _ := firebaseStorage.DefaultBucket()
	w := bucket.Object(filename).NewWriter(c)
	w.ObjectAttrs.ContentType = filetype
	//createImageUrl(imagePath, config.StorageBucket, ctx, client)
	if _, err = io.Copy(w, file); err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	defer file.Close()
	if err := w.Close(); err != nil {
		log.Fatalln(err)
		return
	}

	result := responses.Response{
		Code:    http.StatusOK,
		Message: success.SuccessUpload,
	}

	response.Response(c, &result)
}
