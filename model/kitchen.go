package model

import (
	"time"
)

type (
	// TypesOfFire is a type of fire.
	TypesOfFire string
)

const (
	// TypesOfFireTogether is an order type together
	TypesOfFireTogether TypesOfFire = "together"
	// TypesOfFireAsReady is an order type as_ready
	TypesOfFireAsReady TypesOfFire = "as_ready"
	// TypesOfFireAddOn is an order type add_on
	TypesOfFireAddOn TypesOfFire = "add_on"
	// TypesOfFireRush is an order type rush
	TypesOfFireRush TypesOfFire = "rush"
	// TypesOfFireDont is an order type dont
	TypesOfFireDont TypesOfFire = "dont"
	// TypesOfFireTimedFire is an order type timed_fire
	TypesOfFireTimedFire TypesOfFire = "timed_fire"
	// TypesOfFireByCourse is an order type by_course
	TypesOfFireByCourse TypesOfFire = "by_course"
)

var (
	// TypesOfFireMap converts a TypesOfFire to a printable string
	TypesOfFireMap = map[TypesOfFire]string{
		TypesOfFireTogether:  "Together",
		TypesOfFireAsReady:   "Fire as Ready",
		TypesOfFireAddOn:     "Addon",
		TypesOfFireRush:      "Rush",
		TypesOfFireDont:      "Don't",
		TypesOfFireTimedFire: "Timed Fire",
		TypesOfFireByCourse:  "By Course",
	}
)

// KitchenReceipt kitchen printable model
type KitchenReceipt struct {
	CreatedAt    time.Time
	DineInInfo   *DineInKitchenInfo
	CustomerInfo *CustomerInfo
	OrderType    TypesOfOrder
	Kitchen      string
	Items        []KitchenItem
}

// KitchenItem ordered item printable model
type KitchenItem struct {
	Name       string
	Quantity   int
	Weight     int //to be implemented (name to be decided, could be division)
	FireType   TypesOfFire
	Seats      string
	SubEntries []KitchenSubItem
}

// KitchenSubItem is a modifier for an item
type KitchenSubItem struct {
	Name        string
	Description string
	Index       int
}

// DineInKitchenInfo is the info to impress in the kitchen print
type DineInKitchenInfo struct {
	SectionName string
	Tables      string
	RunnerName  string
}
