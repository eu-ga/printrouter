package model

import (
	"time"

	kModel "github.com/rockspoon/rs.com.kitchen-display-model/model"
)

// KitchenReceiptRequest print kitchen card request
type KitchenReceiptRequest struct {
	Card               kModel.Card `json:"card"`
	Station            string      `json:"station"`
	Language           string      `json:"language"`
	IsPrintedForRunner bool        `json:"isPrintedForRunner"`
	DeliveryAddress    string      `json:"deliveryAddress"`
}

// ToKitchenReceipt convert a print request into a printable model
func (k KitchenReceiptRequest) ToKitchenReceipt() KitchenReceipt {
	items := make([]KitchenItem, 0)
	for i := 0; i < len(k.Card.Items); i++ {
		name := ""
		for _, title := range k.Card.Items[i].Title {
			if title.Language == k.Language {
				name = title.Value
			}
		}
		seats := make([]string, 0)
		for _, diner := range k.Card.Items[i].OrderItemMeta.Diner {
			seats = append(seats, diner.Seat)
		}
		item := KitchenItem{
			Name:       name,
			Quantity:   1,
			SeatNumber: seats,
			IsAllSeats: k.Card.Items[i].OrderItemMeta.IsAllSeats,
			IsSplit:    len(seats) > 1,
			Modifiers:  k.Card.Items[i].OrderItemMeta.ItemMeta.Description,
		}
		items = append(items, item)
	}

	return KitchenReceipt{
		Server:             k.Card.OrderMeta.WaiterName,
		Station:            k.Station,
		FireType:           string(k.Card.OrderMeta.OrderType),
		Timestamp:          k.Card.CreatedAt,
		InvoiceNumber:      k.Card.ID.Hex(),
		TableNumber:        k.Card.OrderMeta.TableNumber,
		Items:              items,
		IsPrintedForRunner: k.IsPrintedForRunner,
		OrderType:          string(k.Card.OrderMeta.Type),
		DeliveryAddress:    k.DeliveryAddress,
	}
}

// KitchenItem ordered item printable model
type KitchenItem struct {
	Name       string
	Quantity   int
	SeatNumber []string
	IsAllSeats bool
	IsSplit    bool
	Modifiers  string
}

// KitchenReceipt kitchen printable model
type KitchenReceipt struct {
	Server             string
	Station            string
	FireType           string
	Timestamp          time.Time
	InvoiceNumber      string
	TableNumber        string
	Items              []KitchenItem
	IsPrintedForRunner bool
	OrderType          string
	DeliveryAddress    string
}
