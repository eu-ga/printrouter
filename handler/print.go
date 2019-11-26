package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	m "github.com/rockspoon/go-common/middleware"
	"github.com/rockspoon/go-common/util"
	mm "github.com/rockspoon/rs.cor.middleware/model"
	s "github.com/rockspoon/rs.cor.middleware/soajs"
	e "github.com/rockspoon/rs.cor.printer-ms/error"
	"github.com/rockspoon/rs.cor.printer-ms/model"
)

type printService interface {
	KitchenReceipt(request model.KitchenReceipt, cData *mm.ContextData) (*model.Payload, error)
	TableBill(request model.Bill, cData *mm.ContextData) (*model.Payload, error)
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

	router.Path("/test").
		Methods(http.MethodGet).HandlerFunc(handler.printTest)

	router.Path("/sales-summary-report").
		Methods(http.MethodGet).HandlerFunc(handler.printSalesSummaryReport)

	router.Path("/payment-invoice").
		Methods(http.MethodGet).HandlerFunc(handler.printPaymentInvoice)

	router.Path("/delivery-takeout/{id}/receipt").
		Methods(http.MethodGet).HandlerFunc(handler.printDeliveryTakeoutReceipt)

	router.Path("/kitchen-card").
		Methods(http.MethodGet).HandlerFunc(handler.printKitchenCard)

	router.Path("/qsr/{id}/receipt").
		Methods(http.MethodGet).HandlerFunc(handler.printQSRReceipt)

	router.Path("/table-bill").
		Methods(http.MethodGet).HandlerFunc(handler.printTableBill)

	router.Path("/venue/{id}/pos-printers").
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

// nolint
type kitchenReceipt struct {
	model.KitchenReceipt
}

// Build builds the create role JSONRequest.
func (r *kitchenReceipt) Build(req *http.Request) error {
	err := json.NewDecoder(req.Body).Decode(r)
	defer util.CloseOrLog(req.Body)
	if err != nil && err != io.EOF {
		return e.DecodeBody(err)
	}
	return nil
}

// Validate validates the create role JSONRequest.
func (r kitchenReceipt) Validate() error {
	return nil
}

// swagger:operation GET /print/kitchen-card Print printKitchenCard
// this endpoint returns a kitchen order ticket
// ---
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Parameters:
//     - name: Receipt Request
//       in: body
//       required: true
//       type: string
//       description: kitchen order data
//
//     Responses:
//       '200':
//         description: printKitchenCard response
//         schema:
//           $ref: "#/definitions/Payload"

func (handler printRouter) printKitchenCard(w http.ResponseWriter, r *http.Request) {
	var req kitchenReceipt
	err := m.ParseRequest(r, &req)
	if err != nil {
		m.JSONError(w, err)
		return
	}
	data, err := s.RequestContextData(r)
	if err != nil {
		m.JSONError(w, err)
		return
	}

	result, err := handler.service.KitchenReceipt(req.KitchenReceipt, data)
	if err != nil {
		m.JSONError(w, err)
		return
	}
	m.JSONReturn(w, http.StatusOK, result)
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

type tableBillRequest struct {
	model.Bill
}

// Build builds the create role JSONRequest.
func (r *tableBillRequest) Build(req *http.Request) error {
	err := json.NewDecoder(req.Body).Decode(r)
	defer util.CloseOrLog(req.Body)
	if err != nil && err != io.EOF {
		return e.DecodeBody(err)
	}
	return nil
}

// Validate validates the create role JSONRequest.
func (r tableBillRequest) Validate() error {
	return nil
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
	var req tableBillRequest
	err := m.ParseRequest(r, &req)
	if err != nil {
		m.JSONError(w, err)
		return
	}
	data, err := s.RequestContextData(r)
	if err != nil {
		m.JSONError(w, err)
		return
	}

	result, err := handler.service.TableBill(req.Bill, data)
	if err != nil {
		m.JSONError(w, err)
		return
	}
	m.JSONReturn(w, http.StatusOK, result)
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
