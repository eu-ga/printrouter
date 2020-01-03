package model

// Printer is the printer response from the get defaul printer endpoint
type Printer struct {
	ID           string   `json:"id,omitempty"`
	Name         string   `json:"name"`
	HardwareName string   `json:"hardwareName"`
	IP           string   `json:"ipAddress"`
	MAC          string   `json:"mac"`
	Model        string   `json:"model"`
	Serial       string   `json:"serialNumber"`
	Topics       []string `json:"topics"`
	Groups       []string `json:"groups"`
}
