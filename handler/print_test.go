package handler

import (
	"net/http"
	"testing"

	"github.com/rockspoon/rs.cor.printer-ms/model"
)

func Test_printTableBill(t *testing.T) {
	tt := []BaseHandlerTest{
		{
			Name:         "empty body",
			Path:         "/table-bill",
			HTTPMethod:   http.MethodPost,
			Req:          "1",
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:       "service error",
			Path:       "/table-bill",
			HTTPMethod: http.MethodPost,
			Req: model.Bill{
				AttendantName: "test",
			},
			ExpectedCode: http.StatusInternalServerError,
		},
		{
			RouteHandler: getDefaultPrinter,
			Name:         "service error",
			Path:         "/table-bill",
			HTTPMethod:   http.MethodPost,
			Req: model.Bill{
				AttendantName: "test",
			},
			ExpectedCode: http.StatusOK,
		},
	}
	ExecHandlerTest(tt, t)
}

func Test_printPaymentReceipt(t *testing.T) {
	tt := []BaseHandlerTest{
		{
			Name:         "empty body",
			Path:         "/payment-invoice",
			HTTPMethod:   http.MethodPost,
			Req:          "1",
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:       "service error",
			Path:       "/payment-invoice",
			HTTPMethod: http.MethodPost,
			Req: model.PaymentReceipt{
				AttendantName: "test",
			},
			ExpectedCode: http.StatusInternalServerError,
		},
		{
			RouteHandler: getDefaultPrinter,
			Name:         "ok",
			Path:         "/payment-invoice",
			HTTPMethod:   http.MethodPost,
			Req: model.Bill{
				AttendantName: "test",
			},
			ExpectedCode: http.StatusOK,
		},
	}
	ExecHandlerTest(tt, t)
}

func Test_printKitchenCard(t *testing.T) {
	tt := []BaseHandlerTest{
		{
			Name:         "empty body",
			Path:         "/kitchen-card",
			HTTPMethod:   http.MethodPost,
			Req:          "1",
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:       "service error",
			Path:       "/kitchen-card",
			HTTPMethod: http.MethodPost,
			Req: model.PaymentReceipt{
				AttendantName: "test",
			},
			ExpectedCode: http.StatusInternalServerError,
		},
		{
			RouteHandler: getDefaultPrinter,
			Name:         "ok",
			Path:         "/kitchen-card",
			HTTPMethod:   http.MethodPost,
			Req: model.Bill{
				AttendantName: "test",
			},
			ExpectedCode: http.StatusOK,
		},
	}
	ExecHandlerTest(tt, t)
}
