package model

import (
	"time"

	"github.com/rockspoon/rs.com.order-model/model"
	money "github.com/rockspoon/rs.cor.common-money"
)

// PaymentReceipt is the receipt to be printed
type PaymentReceipt struct {
	Restaurant    RestaurantInfo  `json:"restaurant"`
	OrderType     model.OrderType `json:"orderType"`
	AttendantName string          `json:"attendantName"`
	CreatedAt     time.Time       `json:"createdAt"`
	Check         Check           `json:"check"`
	Paid          money.SimpleMoney
	Tips          money.SimpleMoney
	PaymentType   string
	Card          *CardInfo
}

// CardInfo is information about the card
type CardInfo struct {
	Type          string // visa, mastercard
	Number        string // 123456*****456
	Authorization string
	Cardholder    string
}
