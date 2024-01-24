package gw

import "fmt"

type router struct {
	handlers map[string]HandleFunc // 每个路由地址，均需要实现对应的处理方法
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandleFunc),
	}
}

func (r *router) addRoute(method string, pattern string, handler HandleFunc) {
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		fmt.Fprintf(c.Writer, "404 NOT FOUND:%s\n", c.Path)
	}
}
