package gw

import (
	"net/http"
)

type HandleFunc func(c *Context)

// Engine 实现了ServeHTTP接口
type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

// GET 方法
func (engine *Engine) GET(pattern string, handler HandleFunc) {
	engine.router.addRoute("GET", pattern, handler)
}

// POST 方法
func (engine *Engine) POST(pattern string, handler HandleFunc) {
	engine.router.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(add string) (err error) {
	return http.ListenAndServe(add, engine)
}

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	c := newContext(writer, request)
	engine.router.handle(c)
}
