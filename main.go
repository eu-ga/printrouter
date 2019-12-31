// Rockspoon Print Microservice
//
// The purpose of this application is to generate payload for printers
//
// Swagger json genereated using https://goswagger.io/generate/spec.html
// swagger generate spec -m -o swagger.json
//
//     Schemes: http, https
//     Host: localhost:6002
//     BasePath: /
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - jwt
// swagger:meta
package main

import (
	"fmt"
	"net/http"

	log "github.com/rockspoon/rs.cor.common-log"
	middleware "github.com/rockspoon/rs.cor.middleware/v2"
	"github.com/rockspoon/rs.cor.printer-ms/controller"
	"github.com/rockspoon/rs.cor.printer-ms/controller/integration"
	"github.com/rockspoon/rs.cor.printer-ms/handler"
)

func main() {
	middlewareHandler, reg, err := middleware.InitMiddleware("soa.json")
	if err != nil {
		log.Panicf("could not initialize SoaJS: %v", err)
	}

	// Integrations
	deviceMS := integration.NewDeviceMS()

	// Services
	service := controller.NewPrintController(deviceMS)

	log.Info("profile service started")
	log.Panic(
		http.ListenAndServe(
			fmt.Sprintf(":%d", reg.ServiceConfiguration.ServicePort),
			handler.New(service, middlewareHandler),
		))
}
