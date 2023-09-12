package zweb

import "net/http"

type HandlerFunc func(c *Context)

type Engine struct {
	router *Router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

// addRoute 添加路由
func (e *Engine) addRoute(method, path string, handler HandlerFunc) {
	e.router.addRoute(method, path, handler)
}

// GET 添加GET请求路由
func (e *Engine) GET(path string, handler HandlerFunc) {
	e.addRoute("GET", path, handler)
}

// POST 添加POST请求路由
func (e *Engine) POST(path string, handler HandlerFunc) {
	e.addRoute("POST", path, handler)
}

// Run 启动服务
func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

// ServeHTTP 实现http.Handler接口
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	e.router.handle(c)
}
