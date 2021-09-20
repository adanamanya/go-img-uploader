package api

import (
	"gogogo/api/handler"
	"gogogo/entity/responses"
	"gogogo/pkg/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	uploadImg := handler.HandleFileupload
	r.POST("/upload", uploadImg)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "nanii!!??",
		})
	})
	errorHandler(r)
}

func errorHandler(r *gin.Engine) {
	r.HandleMethodNotAllowed = true
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, responses.Response{
			Code:    http.StatusMethodNotAllowed,
			Message: errors.ErrNotAllowed.Error(),
		})
		c.Abort()
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, responses.Response{
			Code:    http.StatusNotFound,
			Message: errors.ErrNotFound.Error(),
		})
		c.Abort()
	})
}
