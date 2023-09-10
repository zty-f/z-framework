package zweb

import "net/http"

// HandlerFunc 定义自己的请求处理器 用户通过重写这个方法实现自己的请求处理逻辑
type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// 添加路由 使用路由规则进行匹配 并执行对应的请求处理器
// 路由规则格式为：http method + url pattern -> handler
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router[method+pattern] = handler
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run 启动服务
// 参数 addr 服务地址
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

// ServeHTTP 实现http.Handler接口 具体的接口请求处理入口，然后根据路由规则进行匹配，并执行对应的请求处理器
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler := engine.router[req.Method+req.URL.Path]
	if handler == nil {
		http.NotFound(w, req)
		return
	}
	handler(w, req)
}
