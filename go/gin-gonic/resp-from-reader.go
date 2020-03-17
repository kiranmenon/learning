package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/image", func(c *gin.Context) {
		resp, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || resp.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		c.DataFromReader(http.StatusOK, resp.ContentLength, resp.Header.Get("content-Type"), resp.Body,
			map[string]string{})

	})
	r.Run(":8085")
}
