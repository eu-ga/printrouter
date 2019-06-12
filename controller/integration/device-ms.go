package integration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rockspoon/go-common/util"
	d "github.com/rockspoon/rs.cor.device-model/model"
)

const (
	defaultPrinterPath = "/printer/default?key="
)

// DeviceMS Device Microservice
type DeviceMS struct {
}

// NewDeviceMS creates a new Device Microservice integration point
func NewDeviceMS() DeviceMS {
	return DeviceMS{}
}

// GetDefaultPrinter returns the default printer for a venue
func (DeviceMS) GetDefaultPrinter(path, key string) (*d.Printer, error) {
	uri := "http://" + path + defaultPrinterPath + key
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, fmt.Errorf("bad http request: %v", err)
	}
	req.Header.Set("Cache-Control", "no-cache")
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed on Get Default Printer request: %v", err)
	}
	defer util.CloseOrLog(resp.Body)
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("couldn't read body: %v", err)
	}
	printer := new(d.Printer)
	err = json.Unmarshal(respBody, printer)
	if err != nil {
		return nil, fmt.Errorf("couldn't unmarshal data device microservice: %v", err)
	}
	return printer, nil
}
