package model

import (
	"time"

	"github.com/rockspoon/rs.com.order-model/model"
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
// swagger:model
type KitchenReceipt struct {
	CreatedAt    time.Time          `json:"createdAt"`
	DineInInfo   *DineInKitchenInfo `json:"dinerInfo,omitempty"`
	CustomerInfo *CustomerInfo      `json:"customerInfo,omitempty"`
	OrderType    model.OrderType    `json:"orderInfo"`
	Kitchen      string             `json:"kitchen"`
	Items        []KitchenItem      `json:"items"`
}

// KitchenItem ordered item printable model
// swagger:model
type KitchenItem struct {
	Name       string      `json:"name"`
	Quantity   int         `json:"quantity"`
	Weight     int         `json:"weight"`
	FireType   TypesOfFire `json:"fireType"`
	Seats      string      `json:"seats"`
	SubEntries string      `json:"subEntries"`
}

// DineInKitchenInfo is the info to impress in the kitchen print
// swagger:model
type DineInKitchenInfo struct {
	SectionName string `json:"sectionName"`
	Tables      string `json:"tables"`
	RunnerName  string `json:"runnerName"`
}
