package api

import (
	"gogogo/api/handler"
	"gogogo/entity/responses"
	"gogogo/pkg/errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	auth := os.Getenv("SECRET_TOKEN")
	r.LoadHTMLGlob("templates/*.html") //load folder of static html
	uploadImg := handler.HandleFileupload
	r.POST("/upload", uploadImg) //upload img route
	//route to check if server is running well
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "nanii!!??",
		})
	})
	//route to serve static html
	r.GET("/form", func(c *gin.Context) {
		c.HTML(http.StatusOK, "form.html", auth)
	})
	errorHandler(r)
}

/* error route handling */
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
