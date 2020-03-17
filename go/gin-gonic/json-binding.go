package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/pushData", func(c *gin.Context) {
		var login data
		if err := c.ShouldBind(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error in binding JSON": err.Error()})
			return
		}
		log.Println("Username:", login.UserName)
		log.Println("Password:", login.Password)
		if login.UserName != "Kiran" || login.Password != "Pass" {
			c.JSON(http.StatusUnauthorized, gin.H{"credential error": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "User logged in"})
	})

	r.Run(":8085")
}

type data struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"pass"`
}
