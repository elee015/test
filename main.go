package main

import (
	"fmt"
	"testserver/httpd/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	//init repo
	// repo := model.New()
	r := gin.Default()
	// // Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.GET("/hello", handler.HelloWorld())

	fmt.Println("Spinning server...")

	r.Run(":8080")

}
