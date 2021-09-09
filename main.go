package main
import "net/http"
import "gee/Gee"

func main() {
	r := Gee.New()
	r.GET("/", func(c *Gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *Gee.Context) {
		// expect /hello?name=Geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *Gee.Context) {
		c.JSON(http.StatusOK, Gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}