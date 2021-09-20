package handler

import (
	"fmt"
	"net/http"

	"gogogo/pkg/errors"
	"gogogo/pkg/response"
	"gogogo/pkg/success"

	"gogogo/entity/responses"
	"gogogo/repository"
	"strings"

	"log"

	"github.com/gin-gonic/gin"
)

func HandleFileupload(c *gin.Context) {
	file, err := c.FormFile("image")
	filename := file.Filename
	const MaxFileSize = 8 * 1000000 //8 megabytes (in bytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": errors.ErrParsingFile.Error(),
			"error":   true,
		})
		return
	}
	filetype := file.Header.Get("Content-Type")
	if !strings.Contains(filetype, "image") {
		c.JSON(http.StatusForbidden, gin.H{
			"message": errors.ErrFalseContent.Error(),
			"error":   true,
		})
		return
	}
	fmt.Println("contentsize", file.Size)
	fmt.Println("contentsize", MaxFileSize)
	if file.Size > MaxFileSize {
		c.JSON(http.StatusRequestEntityTooLarge, gin.H{
			"message": errors.ErrTooLarge.Error(),
			"error":   true,
		})
		return
	}
	firebaseStorage := repository.FirebaseStorage()
	bucket, _ := firebaseStorage.DefaultBucket()
	wc := bucket.Object(filename).NewWriter(c)
	wc.ObjectAttrs.ContentType = filetype

	if err := wc.Close(); err != nil {
		log.Fatalln(err)
		return
	}

	result := responses.Response{
		Code:    http.StatusOK,
		Message: success.SuccessUpload,
	}

	response.Response(c, &result)
}
