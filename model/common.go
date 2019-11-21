package model

import (
	"github.com/rockspoon/rs.com.order-model/model"
	"github.com/rockspoon/rs.cor.common-model/address"
)

// DineInOptions location about dinein
type DineInOptions struct {
	SectionName string `json:"sectionName"`
	Tables      string `json:"tables"`
	Seats       string `json:"seats"`
}

// CustomerInfo is info about customere
type CustomerInfo struct {
	Name    string           `json:"name"`
	Address *address.Address `json:"address,omitempty"`
	Phone   string           `json:"phone"`
}

var (
	// TypesOfOrderMap converts a TypesOfOrder to a printable string
	TypesOfOrderMap = map[model.OrderType]string{
		model.OrderTypeDineIn:   "Dine-in",
		model.OrderTypeQSR:      "Takeout",
		model.OrderTypeDelivery: "Delivery",
		model.OrderTypeTakeout:  "Catering",
	}
)
