package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var age int = 10
var Name string = "Wayne"

func test() {
	age := 11
	Name := "Tom"
	fmt.Println("test()", age)
	fmt.Println("test()", Name)
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
	test()
}
