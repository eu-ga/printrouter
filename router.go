//Package debug provides functionality for printing all routes
package printrouter

import (
	"net/http"

	"github.com/gorilla/mux"
)

var format = "%-6s %-30s --> %s (%d handlers)\n"

//SetFormat sets format of the printed string
func SetFormat(f string) {
	format = f
}

//Struct that implements part of the http.Handler interface
//and has wrappers for create handlers
type Router struct {
	*mux.Router
	nuHandlers int
	prefix     string
}

//NewRouter gets pointer to mux.Router and returns new Router that contains mux
func NewRouter(r *mux.Router, prefix string) *Router {
	return &Router{r, 0, prefix}
}

func (r *Router) SetMode(m string) {
	mode = m
}

//Path wrapper for the mux.Router.Path() method
func (r *Router) Path(tpl string) *Route {
	var route = &Route{}

	r.nuHandlers += 1
	route.no = r.nuHandlers

	route.Route = r.NewRoute().Path(tpl)
	route.relativePath = tpl
	route.prefix = r.prefix
	return route
}

//Route struct that implements wrappers for mux.Route structure
type Route struct {
	*mux.Route
	relativePath string
	methods      []string
	f            func(http.ResponseWriter, *http.Request)
	no           int
	prefix       string
}

//Methods wrapper for mux.Route.Method()
func (r *Route) Methods(methods ...string) *Route {
	r.methods = append(r.methods, methods...)
	r.Route = r.Route.Methods(methods...)
	return r
}

//HandlerFunc wrapper for the mux.Route.HandlerFunc() that prints information about http.handlers
func (r *Route) HandlerFunc(f func(http.ResponseWriter, *http.Request)) *mux.Route {
	r.f = f
	handlerName := nameOfFunction(f)
	absolutePath := calculateAbsolutePath(r.prefix, r.relativePath)
	debugPrint(format, r.methods, absolutePath, handlerName, r.no)
	return r.Handler(http.HandlerFunc(f))
}
