package model

import (
	"github.com/rockspoon/rs.cor.common-model/address"
	money "github.com/rockspoon/rs.cor.common-money"
)

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

// Bill is a collection of information that have to be printed in the bill
type Bill struct {
	Restaurant    RestaurantInfo
	OrderType     TypesOfOrder
	AttendantName string // Can be the waiter in dineIn or the person who answered the phone in Delivery and takeout
	Checks        []Check
}

// RestaurantInfo is the info about the restaurant
type RestaurantInfo struct {
	Name    string
	Address address.Address
	Phone   *string
}

// Check to be printed
type Check struct {
	DineInOptions *DineInOptions // Used only on DineIn
	CustomerInfo  *CustomerInfo  // Used only on delivery e takeout

	Items    []EntryItem
	Subtotal money.SimpleMoney // Sum of all items finalPrices + subEntries finalPrices
	Charges  []SubEntry
	Total    money.SimpleMoney // Subtotal + charge
}

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

// EntryItem is an item
type EntryItem struct {
	Name       string
	Quantity   int
	UnityPrice money.SimpleMoney
	FinalPrice money.SimpleMoney // UnityPrice * quantity
	SubEntries []SubEntry
	Weight     int //to be implemented (name to be decided, could be division)
}

// SubEntry is description of an ite
type SubEntry struct {
	Name        string
	Description string
	Index       int               // Order to exhibite the subEntry
	UnityPrice  money.SimpleMoney // Can be zero if only the finalPrice matters
	FinalPrice  money.SimpleMoney // UnityPrice * quantity of entry item
}
