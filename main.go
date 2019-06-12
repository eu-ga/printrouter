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
	"os"

	"github.com/rockspoon/go-common/log"
	"github.com/rockspoon/go-common/soajs"
	"github.com/rockspoon/rs.cor.printer-ms/controller"
	"github.com/rockspoon/rs.cor.printer-ms/controller/integration"
	"github.com/rockspoon/rs.cor.printer-ms/handler"
)

func main() {

	// Integrations
	deviceMS := integration.NewDeviceMS()

	// Services
	service := controller.NewPrintController(deviceMS)

	// Middlewares
	wd, err := os.Getwd()
	if err != nil {
		log.Panicf("could not get working directory: %v", err)
	}
	soajsMiddleware, soaConv, err := soajs.InitSoajs(wd + "/soa.json")
	if err != nil {
		log.Panicf("could not initialize SoaJS: %v", err)
	}

	// Router
	router := handler.New(service, soajsMiddleware)

	log.Info("profile service started")
	log.Panic(http.ListenAndServe(fmt.Sprintf(":%d", soaConv.ServicePort), router))
}
