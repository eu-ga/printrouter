package model

import (
	"time"

	"github.com/rockspoon/rs.cor.common-model/address"
	money "github.com/rockspoon/rs.cor.common-money"
)

// Bill is a collection of information that have to be printed in the bill
type Bill struct {
	Restaurant    RestaurantInfo `json:"restaurant"`
	OrderType     TypesOfOrder   `json:"orderType"`
	AttendantName string         `json:"attendantName"`
	CreatedAt     time.Time      `json:"createdAt"`
	Checks        []Check        `json:"checks"`
}

// RestaurantInfo is the info about the restaurant
type RestaurantInfo struct {
	Name    string          `json:"name"`
	Address address.Address `json:"address"`
	Phone   *string         `json:"phone,omitempty"`
}

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

func (s SubEntrySlice) Len() int {
	return len(s)
}
func (s SubEntrySlice) Less(i, j int) bool {
	return s[i].Index < s[j].Index
}
func (s SubEntrySlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
