package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/gorilla/mux"
	util "github.com/rockspoon/rs.cor.common-util"
	m "github.com/rockspoon/rs.cor.middleware/v2"
	mmocks "github.com/rockspoon/rs.cor.middleware/v2/mocks"
	mmodel "github.com/rockspoon/rs.cor.middleware/v2/model"
	"github.com/rockspoon/rs.cor.printer-ms/controller"
	"github.com/rockspoon/rs.cor.printer-ms/controller/integration"
	dmodel "github.com/rockspoon/rs.cor.printer-ms/controller/integration/model"
	"github.com/rockspoon/rs.cor.printer-ms/dependency"
)

type testHandler struct {
	data     *mmodel.ContextData
	deviceMS dependency.DeviceMS
}

func createTestHandler() (*testHandler, context.Context, *mmodel.ContextData) {
	data := &mmodel.ContextData{}
	data.Venue.ID = util.DBObjectID().Hex()
	data.Profile.ID = util.DBObjectID().Hex()

	ctx := mmocks.ContextFromData(data)
	ctx = mmocks.CreateContext(ctx).
		SetReqID(util.DBObjectID()).
		SetPerfBy(util.DBObjectID().Hex())

	th := new(testHandler)
	th.deviceMS = integration.NewDeviceMS()
	th.data = data

	return th, ctx, data
}
func getDefaultPrinter(sr *SetupResult) map[string]func(http.ResponseWriter, *http.Request) {
	return map[string]func(http.ResponseWriter, *http.Request){
		"/device/printer/default": func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			b, _ := json.Marshal(dmodel.Printer{
				Name: "test",
			})
			_, _ = w.Write(b)
		},
	}
}

// nolint
func (th testHandler) createHTTPExpect(t *testing.T, mw mmodel.Middleware) *httpexpect.Expect {
	service := controller.NewPrintController(th.deviceMS)

	handler := New(service, mw)

	api := mux.NewRouter()
	api.PathPrefix("/").Handler(m.Wrapper(
		handler,
		m.Soajs(mw),
	))

	return httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(api),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	})
}
