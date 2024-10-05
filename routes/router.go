package routes

import (
	"context"
	"net/http"
	"strings"
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
    pathItems := strings.Split(path, "/")
    patternItems := strings.Split(pattern, "/")

    if len(pathItems) != len(patternItems) {
        return nil, false
    }

    params := make(map[string]string)
    for i, pathItem := range pathItems {
        if strings.HasPrefix(patternItems[i], ":") {
            key := strings.Trim(patternItems[i], ":")
            params[key] = pathItem
        } else if pathItem != patternItems[i]{
            return nil, false
        }
    }
    return params, true
}

func contextWithParams(ctx context.Context, params map[string]string) context.Context {
    return context.WithValue(ctx, "params", params)
}

func Params(r *http.Request) map[string]string {
    if params, ok := r.Context().Value("params").(map[string]string); ok {
        return params
    }
    return nil
}


