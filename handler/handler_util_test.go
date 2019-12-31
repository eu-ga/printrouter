package handler

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/gorilla/mux"
	mmocks "github.com/rockspoon/rs.cor.middleware/v2/mocks"
	mmodel "github.com/rockspoon/rs.cor.middleware/v2/model"
	"github.com/rockspoon/rs.cor.printer-ms/utests"
	"github.com/rockspoon/rs.go-test-util/format"
)

var routerMU sync.Mutex

// BaseHandlerTest basic structure of an request validation test -> Do not set properties that are not going to be used
type BaseHandlerTest struct {
	Name                  string
	Req                   interface{}
	Path                  string
	SetupPreTestDBs       func(ctx context.Context, th *testHandler) *SetupResult
	AssertPosTestDBStates func(ctx context.Context, sr *SetupResult, th *testHandler, cData *mmodel.ContextData, res *httpexpect.Response) error
	ExpectedCode          int
	ExpectedBody          string
	HTTPMethod            string
	InvalidContext        bool
	RouteHandler          func(sr *SetupResult) map[string]func(http.ResponseWriter, *http.Request)
	MountRequest          func(sr *SetupResult) interface{}
	MountPath             func(sr *SetupResult) (string, interface{})
}

type SetupResult struct {
}

type singleHandlerTest struct {
	BaseHandlerTest
	utests.TestMock
}
type internalHT struct {
	base BaseHandlerTest
	t    *testing.T
	th   *testHandler
	ctx  context.Context
	data *mmodel.ContextData
}

// GetTestName return test name
func (inst internalHT) GetTestName() string {
	return inst.base.Name
}

// ExecuteTest execute test and return an eventual error
func (inst internalHT) ExecuteTest() error {
	return handlerTest(inst.ctx, inst.base, inst.t, inst.th, inst.data)
}

func handlerTest(ctx context.Context, tc BaseHandlerTest, t *testing.T, th *testHandler, data *mmodel.ContextData) error {
	// Setup environment
	var testData *mmodel.ContextData
	if tc.InvalidContext {
		testData = nil
	} else {
		testData = data
	}

	var setupResult *SetupResult
	if tc.SetupPreTestDBs != nil {
		setupResult = tc.SetupPreTestDBs(ctx, th)
	}

	// HTTP setup
	// Create router to handle outside calls
	router := mux.NewRouter()
	if tc.RouteHandler != nil {
		routesHandlers := tc.RouteHandler(setupResult)
		for outPath := range routesHandlers {
			router.HandleFunc(outPath, routesHandlers[outPath])
		}
	}
	routerMU.Lock()
	ts := httptest.NewServer(router)

	// Makes sure the clients are disconnected before trying to close the server so it does not get stuck waiting for ghost requests
	defer func() { ts.Close(); routerMU.Unlock() }()

	baseURL := strings.Split(ts.URL, ":")
	port, _ := strconv.Atoi(baseURL[2])
	paths := make(map[string]string)
	paths[mmodel.ImageMS] = "http://localhost:" + baseURL[2] + "/image/"
	paths[mmodel.DEVICE] = "http://localhost:" + baseURL[2] + "/device/"
	headers := map[string]string{
		"key":          "key",
		"access_token": "token",
	}
	data.Paths = paths
	mw := mmocks.Middleware(testData, headers, baseURL[0]+":"+baseURL[1], port)

	expect := th.createHTTPExpect(t, mw)
	var httpFunc func(path string, pathargs ...interface{}) *httpexpect.Request
	switch tc.HTTPMethod {
	case "POST":
		httpFunc = expect.POST
	case "GET":
		httpFunc = expect.GET
	case "PUT":
		httpFunc = expect.PUT
	case "DELETE":
		httpFunc = expect.DELETE
	case "PATCH":
		httpFunc = expect.PATCH
	default:
		log.Panic("No HTTPMethod passed -> Please Specify the method to be called")
	}
	if tc.Path == "" && tc.MountPath == nil {
		log.Panic("No Path or MountPath given, no way to know where the request should go")
	}

	// Use or Mount Request
	var request interface{}
	if tc.MountRequest != nil {
		request = tc.MountRequest(setupResult)
	} else {
		request = tc.Req
	}
	// Use or Mount Path
	var path string
	var query interface{}
	if tc.MountPath != nil {
		path, query = tc.MountPath(setupResult)
	} else {
		path = tc.Path
	}
	// Mount request
	preRequest := httpFunc(path).
		WithText(format.ToStringJSON(request))
	if query != nil {
		preRequest = preRequest.WithQueryObject(query)
	}
	// Send
	resp := preRequest.Expect().
		Status(tc.ExpectedCode)
	if tc.ExpectedBody != "" {
		resp.Body().Equal(tc.ExpectedBody)
	}
	if tc.AssertPosTestDBStates != nil {
		err := tc.AssertPosTestDBStates(ctx, setupResult, th, data, resp)
		return err
	}
	return nil
}
func transformSingleBaseToHandlerTest(base BaseHandlerTest, t *testing.T) singleHandlerTest {
	th, ctx, data := createTestHandler()
	tMock := internalHT{base, t, th, ctx, data}
	valTest := singleHandlerTest{
		base, tMock,
	}
	return valTest
}

func transformAllBaseToHandlerTest(base []BaseHandlerTest, t *testing.T) []utests.TestMock {
	handlerTests := []utests.TestMock{}
	for i := range base {
		handlerTests = append(handlerTests, transformSingleBaseToHandlerTest(base[i], t))
	}
	return handlerTests
}

// ExecHandlerTest run handler tests between the expected request format and the mock input
func ExecHandlerTest(base []BaseHandlerTest, t *testing.T) []error {
	valt := transformAllBaseToHandlerTest(base, t)
	return utests.RunTest(valt, t)
}
