package debug

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var format = "%-6s %-25s --> %s (%d handlers)\n"

//SetFormat sets format of the printed string
func SetFormat(f string) {
	format = f
}

type Router struct {
	*mux.Router
	nuHandlers int
}

func NewRouter(r *mux.Router) *Router {
	return &Router{r, 0}
}

func (r *Router) Path(tpl string) *Route {
	fmt.Println("Route path: ", tpl)
	var route = &Route{
		no: 1,
	}

	r.nuHandlers = route.no

	route.Route = r.NewRoute().Path(tpl)
	route.relativePath = tpl
	return route
}

type Route struct {
	*mux.Route
	relativePath string
	methods      []string
	f            func(http.ResponseWriter, *http.Request)
	no           int
}

func (r *Route) Methods(methods ...string) *Route {
	r.methods = append(r.methods, methods...)
	return r.Methods(methods...)
}

func (r *Route) HandlerFunc(f func(http.ResponseWriter, *http.Request)) *mux.Route {
	r.f = f
	handlerName := nameOfFunction(f)
	absolutePath := calculateAbsolutePath(r.relativePath)
	debugPrint(format, r.methods, absolutePath, handlerName, r.no)
	return r.Handler(http.HandlerFunc(f))
}
