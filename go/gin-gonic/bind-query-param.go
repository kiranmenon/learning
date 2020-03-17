package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name     string `form:"name"`
	LastName string `form:"lname"`
}

func main() {
	r := gin.Default()
	r.GET("/test", getFn)
	r.Run(":8085")
}

func getFn(c *gin.Context) {
	var p Person
	if c.ShouldBind(&p) == nil {
		log.Println("Person name: ", p.Name)
		log.Println("Last name: ", p.LastName)
	}
	c.String(http.StatusOK, "success")
}
