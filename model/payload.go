package model

// Payload printer payload wrapper
// swagger:model
type Payload struct {
	PrintPayload    string `json:"printPayload"`
	IPAddress       string `json:"ipAddress"`
	PrinterModel    string `json:"printerModel"`
	DescribeMessage string `json:"describeMessage"`
	ID              string `json:"id"`
}

// VenuePrinterPayload venue printer payload
// swagger:model
type VenuePrinterPayload struct {
	DeviceID         string            `json:"deviceId"`
	Type             string            `json:"type"`
	Status           string            `json:"status"`
	UID              string            `json:"uid"`
	Name             string            `json:"name"`
	Version          string            `json:"version"`
	PreferredNetwork string            `json:"preferredNetwork"`
	Metadata         map[string]string `json:"metadata"`
}
