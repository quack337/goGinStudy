package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

type Person struct {
  Name string  `form:"name"`
  Age  int     `form:"age"`
  Enabled bool `form:"enabled"`
  Color string `form:"color"`
  Num int      `form:"num"`
}

func main() {
  router := gin.Default()
  router.SetFuncMap(template.FuncMap{
	"checked": func (value bool) string {
		if value { return "checked" }
		return ""
	},
	"selected": func (value bool) string {
		if value { return "selected" }
		return ""
	},
  })
  router.LoadHTMLGlob("templates/*")

  router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "form1.html", gin.H{
      "person": Person{Name: "홍길동", Age: 16},
    })
  })
  router.POST("/", func(c *gin.Context) {
	var person Person
	c.Bind(&person)
	fmt.Println(person)
	person.Age++
    c.HTML(http.StatusOK, "form1.html", gin.H{
      "person": person,
    })
  })

  router.Run()
}
