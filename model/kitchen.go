package model

import (
	"time"
)

type KitchenItem struct {
	Name       string
	Quantity   int64
	SeatNumber []string
	IsAllSeats bool
	IsSplit    bool
	Modifiers  string
}

type KitchenReceipt struct {
	Server             string
	Station            string
	FireType           string
	Timestamp          time.Time
	InvoiceNumber      int64
	TableNumber        string
	Items              []KitchenItem
	IsPrintedForRunner bool
	OrderType          string
	DeliveryAddress    string
}
