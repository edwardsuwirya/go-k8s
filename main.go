package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	port, isPortExist := os.LookupEnv("PORT")
	response, isResponseExist := os.LookupEnv("RESPONSE")
	if !isPortExist {
		port = "8080"
	}
	if !isResponseExist {
		response = "pong"
	}
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "My Hello World",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": response + "!",
		})
	})
	err := r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
}
