package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	h "github.com/rockspoon/go-common/handler"
	m "github.com/rockspoon/go-common/middleware"
)

type (
	// Router represents gorilla mux driver.
	Router struct {
		*mux.Router
	}
)

// New creates a router for this microservice.
func New(
	service printService,
	middleware func(http.Handler) http.Handler,
) Router {
	router := Router{
		Router: mux.NewRouter(),
	}
	router.HandleFunc("/healthz", h.HealthzHandler)

	router.PathPrefix("/").Handler(m.Wrapper(
		newPrintRouter(service),
		m.Soajs(middleware),
	))

	return router
}
