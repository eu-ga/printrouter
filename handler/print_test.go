package handler

import (
	"errors"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/gorilla/mux"
	"github.com/rockspoon/go-common/middleware"
	s "github.com/rockspoon/rs.cor.middleware/soajs"
	"github.com/rockspoon/rs.cor.printer-ms/model"
)

type ControllerMock struct {
	Payload *model.Payload
	Error   error
}

func (c ControllerMock) KitchenReceipt(request model.KitchenReceiptRequest, cData *s.ContextData) (*model.Payload, error) {
	if c.Error != nil {
		return nil, c.Error
	}
	return c.Payload, nil
}

func TestRouter_KitchenReceipt(t *testing.T) {

	tt := []struct {
		name            string
		soaMiddleware   middleware.Middleware
		request         string
		payload         *model.Payload
		controllerError error
		contextData     *s.ContextData
		expectedCode    int
		expectedBody    string
	}{
		{
			name:          "bad request",
			request:       "not a json",
			soaMiddleware: soajsTest(false, "", "", ""),
			expectedCode:  http.StatusBadRequest,
			expectedBody:  `{"code":"BadBodyFormat","error":"could not decode request body","cause":"invalid character 'o' in literal null (expecting 'u')"}`,
		},
		{
			name:          "no soa context",
			request:       `{}`,
			soaMiddleware: soajsTest(false, "", "", ""),
			expectedCode:  http.StatusInternalServerError,
			expectedBody:  `{"code":"UnexpectedError","error":"could not type assert soajs context"}`,
		},
		{
			name:            "controller error",
			request:         `{}`,
			soaMiddleware:   soajsTest(true, "", "", ""),
			contextData:     getContextData("", "", ""),
			controllerError: errors.New("could not find printer for venue"),
			expectedCode:    http.StatusInternalServerError,
			expectedBody:    `{"code":"UnexpectedError","error":"could not find printer for venue"}`,
		},
		{
			name:          "success",
			request:       `{}`,
			soaMiddleware: soajsTest(true, "1", "1", "1"),
			contextData:   getContextData("1", "1", "1"),
			payload:       new(model.Payload),
			expectedCode:  http.StatusOK,
			expectedBody:  `{"printPayload":"","ipAddress":"","printerModel":"","describeMessage":""}`,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			controller := ControllerMock{Payload: tc.payload, Error: tc.controllerError}
			handler := New(controller, tc.soaMiddleware)

			api := mux.NewRouter()
			api.PathPrefix("/").Handler(middleware.Wrapper(
				handler,
				middleware.Soajs(tc.soaMiddleware),
			))

			e := httpexpect.WithConfig(httpexpect.Config{
				Client: &http.Client{
					Transport: httpexpect.NewBinder(api),
					Jar:       httpexpect.NewJar(),
				},
				Reporter: httpexpect.NewAssertReporter(t),
			})
			e.GET("/kitchen-order").
				WithText(tc.request).
				Expect().
				Status(tc.expectedCode).
				Body().Equal(tc.expectedBody + "\n")

		})
	}
}
