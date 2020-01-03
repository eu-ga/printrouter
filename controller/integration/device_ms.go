package integration

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	request "github.com/rockspoon/rs.cor.common-request"
	m "github.com/rockspoon/rs.cor.middleware/v2"
	mmodel "github.com/rockspoon/rs.cor.middleware/v2/model"
	d "github.com/rockspoon/rs.cor.printer-ms/controller/integration/model"
	"github.com/rockspoon/rs.cor.printer-ms/errors"
)

const (
	defaultPrinterPath = "%sprinter/default"
)

// DeviceMS Device Microservice
type DeviceMS struct{}

// NewDeviceMS creates a new Device Microservice integration point
func NewDeviceMS() DeviceMS {
	return DeviceMS{}
}

// GetDefaultPrinter returns the default printer for a venue
func (DeviceMS) GetDefaultPrinter(ctx context.Context) (*d.Printer, error) {
	cData, err := m.DataFromContext(ctx)
	if err != nil {
		return nil, errors.InvalidMiddlewareContext()
	}

	devicePath := cData.Paths[mmodel.DEVICE]

	req, err := request.NewRequestWithContext(ctx, "GET", fmt.Sprintf(defaultPrinterPath, devicePath), nil)
	if err != nil {
		return nil, err
	}

	var result = new(d.Printer)
	if statusCode, err := request.DoJSON(req, result); err != nil {
		return nil, errors.InvalidMicroserviceResponse("Print", strconv.Itoa(statusCode))
	} else if statusCode != http.StatusOK {
		return nil, errors.InvalidMicroserviceResponse("Print", strconv.Itoa(statusCode))
	}
	return result, nil
}
