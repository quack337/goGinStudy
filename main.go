package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "form1.html", gin.H{
			"person": Person{Name: "홍길동", Age: 16},
		})
	})
	router.Run()
}
