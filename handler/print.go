package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	m "github.com/rockspoon/go-common/middleware"
	"github.com/rockspoon/rs.cor.printer-ms/model"
)

type printService interface {
}

type (
	printRouter struct {
		*mux.Router
		service printService
	}
)

func newPrintRouter(service printService) printRouter {
	router := mux.NewRouter().PathPrefix("/").Subrouter()
	handler := printRouter{
		router,
		service,
	}

	router.Path("test").
		Methods(http.MethodGet).HandlerFunc(handler.printTest)

	router.Path("sales-summary-report").
		Methods(http.MethodGet).HandlerFunc(handler.printSalesSummaryReport)

	router.Path("payment-invoice").
		Methods(http.MethodGet).HandlerFunc(handler.printPaymentInvoice)

	router.Path("delivery-takeout/{id}/receipt").
		Methods(http.MethodGet).HandlerFunc(handler.printDeliveryTakeoutReceipt)

	router.Path("kitchen-order").
		Methods(http.MethodGet).HandlerFunc(handler.printKitchenOrder)

	router.Path("qsr/{id}/receipt").
		Methods(http.MethodGet).HandlerFunc(handler.printQSRReceipt)

	router.Path("table-bill").
		Methods(http.MethodGet).HandlerFunc(handler.printTableBill)

	router.Path("venue/{id}/pos-printers").
		Methods(http.MethodGet).HandlerFunc(handler.printVenuePrinters)

	return handler
}

// swagger:operation GET /print/test Print printTest
// this endpoint returns a test payload
// ---
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Parameters:
//     - name: deviceId
//       in: query
//       required: true
//       type: string
//       description: printer id
//     - name: firstSetup
//       in: query
//       required: true
//       type: boolean
//       description: flag for first time setup
//
//     Responses:
//       '200':
//         description: printTest response
//         schema:
//           $ref: "#/definitions/Payload"
//
func (handler printRouter) printTest(w http.ResponseWriter, r *http.Request) {
	m.JSONReturn(w, http.StatusOK, model.Payload{})
}

// swagger:operation GET /print/sales-summary-report Print printSalesSummaryReport
// this endpoint returns a a sales summary report payload
// ---
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Parameters:
//     - name: startDate
//       in: query
//       required: true
//       type: string
//       description: start date
//     - name: endDate
//       in: query
//       required: true
//       type: string
//       description: end date
//
//     Responses:
//       '200':
//         description: printSalesSummaryReport response
//         schema:
//           $ref: "#/definitions/Payload"
//
func (handler printRouter) printSalesSummaryReport(w http.ResponseWriter, r *http.Request) {
	m.JSONReturn(w, http.StatusOK, model.Payload{})
}

// swagger:operation GET /print/payment-invoice Print printPaymentInvoice
// this endpoint returns a payment invoice payload by payment id
// ---
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Parameters:
//     - name: userPaymentId
//       in: query
//       required: true
//       type: string
//       description: payment id
//
//     Responses:
//       '200':
//         description: printPaymentInvoice response
//         schema:
//           $ref: "#/definitions/Payload"
//
func (handler printRouter) printPaymentInvoice(w http.ResponseWriter, r *http.Request) {
}

// swagger:operation GET /print/delivery-takeout/{id}/receipt Print printDeliveryTakeoutReceipt
// this endpoint returns a delivery and takeout receipt payload
// ---
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Parameters:
//     - name: id
//       in: path
//       required: true
//       type: string
//       description: dinning party id
//
//     Responses:
//       '200':
//         description: printDeliveryTakeoutReceipt response
//         schema:
//           $ref: "#/definitions/Payload"
//
func (handler printRouter) printDeliveryTakeoutReceipt(w http.ResponseWriter, r *http.Request) {
}

// swagger:operation GET /print/kitchen-order Print printKitchenOrder
// this endpoint returns a kitchen order ticket
// ---
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Parameters:
//     - name: stationName
//       in: query
//       required: true
//       type: string
//       description: preparation station's name
//     - name: itemsFiredIds
//       in: query
//       required: true
//       type: string
//       description: serialized items ids [string1,string2,string3]
//
//     Responses:
//       '200':
//         description: printKitchenOrder response
//         schema:
//           $ref: "#/definitions/Payload"
//
func (handler printRouter) printKitchenOrder(w http.ResponseWriter, r *http.Request) {
}

// swagger:operation GET /print/qsr/{id}/receipt Print printQSRReceipt
// this endpoint returns a qsr receipt payload
// ---
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Parameters:
//     - name: id
//       in: path
//       required: true
//       type: string
//       description: dinning party id
//
//     Responses:
//       '200':
//         description: printQSRReceipt response
//         schema:
//           $ref: "#/definitions/Payload"
//
func (handler printRouter) printQSRReceipt(w http.ResponseWriter, r *http.Request) {
}

// swagger:operation GET /print/table-bill Print printTableBill
// this endpoint returns a table's bill payload
// ---
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Parameters:
//     - name: id
//       in: query
//       required: true
//       type: string
//       description: dinning party id
//
//     Responses:
//       '200':
//         description: printTableBill response
//         schema:
//           $ref: "#/definitions/Payload"
//
func (handler printRouter) printTableBill(w http.ResponseWriter, r *http.Request) {
}

// swagger:operation GET /print/venue/{id}/pos-printers Print printVenuePrinters
// this endpoint returns a printer's information for a venue
// ---
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Parameters:
//     - name: id
//       in: path
//       required: true
//       type: string
//       description: venue id
//
//     Responses:
//       '200':
//         description: printQSRReceipt response
//         schema:
//           $ref: "#/definitions/VenuePrinterPayload"
//
func (handler printRouter) printVenuePrinters(w http.ResponseWriter, r *http.Request) {
}