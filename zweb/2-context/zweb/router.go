package zweb

import "log"

type Router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *Router {
	return &Router{handlers: make(map[string]HandlerFunc)}
}

func (r *Router) addRoute(method string, path string, handler HandlerFunc) {
	log.Printf("Add route %s %s", method, path)
	r.handlers[method+"-"+path] = handler
}

func (r *Router) handle(c *Context) {
	handler, ok := r.handlers[c.Method+"-"+c.Path]
	if !ok {
		c.String(404, "Not Found the Path %s\n", c.Path)
		return
	}
	handler(c)
}
