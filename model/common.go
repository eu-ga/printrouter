package model

import "github.com/rockspoon/rs.cor.common-model/address"

// TypesOfOrder is the type of order
type TypesOfOrder string

const (
	// TypesOfOrderDinein represents dinein available sales option
	TypesOfOrderDinein TypesOfOrder = "dinein"
	// TypesOfOrderTakeout represents takeout available sales option
	TypesOfOrderTakeout TypesOfOrder = "takeout"
	// TypesOfOrderDelivery represents delivery available sales option
	TypesOfOrderDelivery TypesOfOrder = "delivery"
	// TypesOfOrderCatering represents catering available sales option
	TypesOfOrderCatering TypesOfOrder = "catering"
)

// DineInOptions location about dinein
type DineInOptions struct {
	SectionName string
	Tables      string
	Seats       string
}

// CustomerInfo is info about customere
type CustomerInfo struct {
	Name    string
	Address *address.Address // Used only on delivery
	Phone   string
}

var (
	// TypesOfOrderMap converts a TypesOfOrder to a printable string
	TypesOfOrderMap = map[TypesOfOrder]string{
		TypesOfOrderDinein:   "Dine-in",
		TypesOfOrderTakeout:  "Takeout",
		TypesOfOrderDelivery: "Delivery",
		TypesOfOrderCatering: "Catering",
	}
)
