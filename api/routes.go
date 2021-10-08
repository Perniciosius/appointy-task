package api

import (
	"context"
	"net/http"
	"regexp"
	"strings"
)

// Custom handler for routes
type RouteHandler struct{}

func (handler RouteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var allow []string
	for _, route := range routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}
			ctx := context.WithValue(r.Context(), ctxKey{}, matches[1:])
			route.handler(w, r.WithContext(ctx))
			return
		}
	}
	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.NotFound(w, r)
}

// Custom route for parsing url parameters
type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

var routes = []route{
	newRoute("POST", "/users", CreateUser),
	newRoute("GET", "/users/([^/]+)", GetUser),
	newRoute("POST", "/posts", CreatePost),
	newRoute("GET", "/posts/([^/]+)", GetPost),
	newRoute("GET", "/posts/users/([^/]+)", GetPostList),
}

func newRoute(method, pattern string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

type ctxKey struct{}

func getField(r *http.Request, index int) string {
	fields := r.Context().Value(ctxKey{}).([]string)
	return fields[index]
}
