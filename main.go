package main

import (
	"gweb/gw"
	"log"
	"net/http"
)

func main() {
	r := gw.New()

	r.GET("/", func(c *gw.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gw</h1>")
	})

	r.GET("/hello", func(c *gw.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gw.Context) {
		c.Json(http.StatusOK, gw.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	log.Println("Run at 9999")
	r.Run(":9999")
}
