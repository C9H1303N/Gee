package main

import (
	"gee/Gee"
	"log"
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
	v := r.Group("/abc")
	v.Static("/assets", "C:/Users/yaozijian/GolandProjects/gee/usr")
	// 或相对路径 r.Static("/assets", "./static")
	r.Run(":9999")

}
