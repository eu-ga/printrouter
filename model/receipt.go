package model

import (
	"time"

	"github.com/rockspoon/rs.com.order-model/model"
	money "github.com/rockspoon/rs.cor.common-money"
)

// PaymentReceipt is the receipt to be printed
// swagger:model
type PaymentReceipt struct {
	Restaurant    RestaurantInfo    `json:"restaurant"`
	OrderType     model.OrderType   `json:"orderType"`
	AttendantName string            `json:"attendantName"`
	CreatedAt     time.Time         `json:"createdAt"`
	Check         Check             `json:"check"`
	Paid          money.SimpleMoney `json:"paid"`
	Tips          money.SimpleMoney `json:"tips"`
	PaymentType   string            `json:"paymentType"`
	Card          *CardInfo         `json:"card"`
}

// CardInfo is information about the card
// swagger:model
type CardInfo struct {
	Type          string `json:"type"`   // visa, mastercard
	Number        string `json:"number"` // 123456*****456
	Authorization string `json:"authorization"`
	Cardholder    string `json:"cardHolder"`
}
