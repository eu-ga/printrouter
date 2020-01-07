package model

// TestPayload is the test to be printed
type TestPayload struct {
	Restaurant RestaurantInfo `json:"restaurant"`
}

// TestBody is the test body request
// swagger:model
type TestBody struct {
	IPAddress  string `json:"ipAddress"`
	PrintModel string `json:"printModel"`
}
