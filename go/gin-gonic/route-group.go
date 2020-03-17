package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "This is GET operation V1.")
		})

		v1.POST("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "This is POST operation V1")
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "This is GET operation V2.")
		})

		v2.POST("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "This is POST operation V2")
		})
	}

	r.Run(":8085")
}
