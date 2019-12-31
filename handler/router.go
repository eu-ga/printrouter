package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	m "github.com/rockspoon/rs.cor.middleware/v2"
	h "github.com/rockspoon/rs.cor.middleware/v2/handler"
	"github.com/rockspoon/rs.cor.printer-ms/dependency"
)

type (
	// Router represents gorilla mux driver.
	Router struct {
		*mux.Router
	}
)

// New creates a router for this microservice.
func New(
	service dependency.PrintService,
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
