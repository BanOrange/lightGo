package main

import (
	"net/http"
	"lightGin"
)


func main() {
	r := lightGin.New()
	r.GET("/index", func(c *lightGin.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *lightGin.Context) {
			c.HTML(http.StatusOK, "<h1>Hello lightGin</h1>")
		})

		v1.GET("/hello", func(c *lightGin.Context) {
			// expect /hello?name=lightGinktutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *lightGin.Context) {
			// expect /hello/lightGinktutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *lightGin.Context) {
			c.JSON(http.StatusOK, lightGin.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":9999")
}