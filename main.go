package main

import (
	"gee/Gee"
	"log"
	"net/http"
	"time"
)

func onlyForV2() Gee.HandlerFunc {
	return func(c *Gee.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		//	c.Fail(500, "Internal Server Error")
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := Gee.New()
	r.Use(Gee.Logger()) // global midlleware
	r.GET("/", func(c *Gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *Gee.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
