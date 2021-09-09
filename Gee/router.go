package Gee

import (
	"log"
	"net/http"
)

type router struct {
	roots map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
		roots: make(map[string]*node),
	}
}

func parsePattern(pattern string) []string {
	
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc)  {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context)  {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok{
		handler(c)
		log.Printf("Request %s url: %s", c.Method, c.Path)
	} else {
		log.Printf("NOT FOUND: %s\n",c.Path)
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n",c.Path)
	}
}