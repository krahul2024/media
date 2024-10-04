package routes

import (
	"context"
	"net/http"
)

type Router struct {
    routes []*route
    prefix string
}

type route struct {
    pattern string
    handler http.HandlerFunc
}

func NewRouter() *Router {
    return &Router{}
}

func (r *Router) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
    r.routes = append(r.routes, &route{
        pattern: r.prefix + pattern,
        handler: handler,
    })
}

func (r *Router) Group(prefix string) *Router {
    return &Router{ routes: r.routes, prefix: r.prefix + prefix }
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    for _, route := range r.routes {
        if params, ok := match(route.pattern, req.URL.Path); ok {
            ctx := contextWithParams(req.Context(), params)
            route.handler(w, req.WithContext(ctx))
            return
        }
    }
    http.NotFound(w, req)
}

func match(pattern, path string) (map[string]string, bool) {
    
}

func contextWithParams(ctx context.Context, parms map[string]string) context.Context {

}


