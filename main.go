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

	r.GET("/hello/:name", func(c *lightGin.Context) {
		// expect /hello/lightGinktutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *lightGin.Context) {
		c.JSON(http.StatusOK, lightGin.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}