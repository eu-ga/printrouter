package model

import (
	"time"

	"github.com/rockspoon/rs.cor.common-model/address"
	money "github.com/rockspoon/rs.cor.common-money"
)

// Bill is a collection of information that have to be printed in the bill
type Bill struct {
	Restaurant    RestaurantInfo
	OrderType     TypesOfOrder
	AttendantName string // Can be the waiter in dineIn or the person who answered the phone in Delivery and takeout
	CreatedAt     time.Time
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
	Charges  SubEntrySlice
	Total    money.SimpleMoney // Subtotal + charge
}

// EntryItem an item in the bill
type EntryItem struct {
	Name       string
	Quantity   int
	UnityPrice money.SimpleMoney
	FinalPrice money.SimpleMoney // UnityPrice * quantity
	SubEntries SubEntrySlice
	Weight     int // TODO to be implemented (name to be decided, could be division)
}

// SubEntry is description of an item
type SubEntry struct {
	Name        string
	Description string
	Index       int               // Order to exhibite the subEntry
	UnityPrice  money.SimpleMoney // Can be zero if only the finalPrice matters
	FinalPrice  money.SimpleMoney // UnityPrice * quantity of entry item
}

// SubEntrySlice implements sort.Interface based on Index
type SubEntrySlice []SubEntry

func (s SubEntrySlice) Len() int {
	return len(s)
}
func (s SubEntrySlice) Less(i, j int) bool {
	return s[i].Index < s[j].Index
}
func (s SubEntrySlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
