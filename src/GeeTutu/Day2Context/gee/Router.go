package gee

import "log"

// 路由类
type router struct {
	handles map[string]HandlerFunc // 路由表
}

func newRouter() *router {
	return &router{
		handles: make(map[string]HandlerFunc),
	}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handles[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handle, ok := r.handles[key]; ok {
		handle(c)
	} else {
		c.String(404, "404 NOT FOUND: %s\n", c.Path)
	}
}
