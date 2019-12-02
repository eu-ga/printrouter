package model

import (
	"github.com/rockspoon/rs.com.order-model/model"
	"github.com/rockspoon/rs.cor.common-model/address"
	money "github.com/rockspoon/rs.cor.common-money"
)

// RestaurantInfo is the info about the restaurant
type RestaurantInfo struct {
	Name    string          `json:"name"`
	Address address.Address `json:"address"`
	Phone   *string         `json:"phone,omitempty"`
}

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

// Check to be printed
type Check struct {
	DineInOptions *DineInOptions `json:"dineInOptions,omitempty"`
	CustomerInfo  *CustomerInfo  `json:"customerInfo,omitempty"`

	Items    []EntryItem       `json:"items"`
	Subtotal money.SimpleMoney `json:"subtotal"`
	Charges  SubEntrySlice     `json:"charges"`
	Total    money.SimpleMoney `json:"total"`
}

// EntryItem an item in the bill
type EntryItem struct {
	Name       string            `json:"name"`
	Quantity   int               `json:"quantity"`
	UnityPrice money.SimpleMoney `json:"unityPrice"`
	FinalPrice money.SimpleMoney `json:"finalPrice"`
	SubEntries SubEntrySlice     `json:"subEntries"`
	Weight     int               `json:"weight"`
}

// SubEntry is description of an item
type SubEntry struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Index       int               `json:"index"`
	UnityPrice  money.SimpleMoney `json:"unityPrice"`
	FinalPrice  money.SimpleMoney `json:"finalPrice"`
}

// SubEntrySlice implements sort.Interface based on Index
type SubEntrySlice []SubEntry

// Len returns the len of SubEntrySlice, this is required to implement the sort.Interface
func (s SubEntrySlice) Len() int {
	return len(s)
}

// Less returns if an element opf the slice is smaller than the other, this is required to implement the sort.Interface
func (s SubEntrySlice) Less(i, j int) bool {
	return s[i].Index < s[j].Index
}

// Swap swaps two elements of a SubEntrySlice, this is required to implement the sort.Interface
func (s SubEntrySlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
