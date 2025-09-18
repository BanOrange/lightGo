package main

import (
	"net/http"
	"lightGin"
)


func main() {
	r := lightGin.New()
	r.GET("/", func(c *lightGin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello lightGin</h1>")
	})
	r.GET("/hello", func(c *lightGin.Context) {
		// expect /hello?name=lightGinktutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *lightGin.Context) {
		c.JSON(http.StatusOK, lightGin.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}