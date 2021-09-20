package handler

import (
	"net/http"
	"os"

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
	//handling form upload
	file, handler, err := c.Request.FormFile("data")
	auth := c.Request.FormValue("auth")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": errors.ErrParsingFile.Error(),
			"error":   true,
		})
		return
	}
	if auth != os.Getenv("SECRET_TOKEN") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": errors.ErrUnauthorized.Error(),
			"error":   true,
		})
		return
	}
	defer file.Close()

	filename := handler.Filename                   //get file name
	const MaxFileSize = 8 * 1000000                //8 megabytes (in bytes)
	filetype := handler.Header.Get("Content-Type") //get content type

	//if content type is not image, will throw error false content
	if !strings.Contains(filetype, "image") {
		c.JSON(http.StatusForbidden, gin.H{
			"message": errors.ErrFalseContent.Error(),
			"error":   true,
		})
		return
	}
	//if size more than 8mb will throw error content too large
	if handler.Size > MaxFileSize {
		c.JSON(http.StatusRequestEntityTooLarge, gin.H{
			"message": errors.ErrTooLarge.Error(),
			"error":   true,
		})
		return
	}
	firebaseStorage := repository.FirebaseStorage()
	bucket, _ := firebaseStorage.DefaultBucket()
	//Write image to firebase storage
	w := bucket.Object(filename).NewWriter(c)
	w.ObjectAttrs.ContentType = filetype //define contenttype

	if _, err = io.Copy(w, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": errors.ErrFailedUpload.Error(),
			"error":   true,
		})
		return
	}
	defer file.Close()
	if err := w.Close(); err != nil {
		log.Fatalln(err)
		return
	}
	//return success
	result := responses.Response{
		Code:    http.StatusOK,
		Message: success.SuccessUpload,
	}

	response.Response(c, &result)
}
