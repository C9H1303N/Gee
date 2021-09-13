package main

import (
	"net/http"

	"gee/Gee"
)

func main() {
	r := Gee.Default()
	r.GET("/", func(c *Gee.Context) {
		c.String(http.StatusOK, "Hello Geektutu\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *Gee.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
