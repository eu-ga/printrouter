package model

import (
	"time"

	order "github.com/rockspoon/rs.com.order-model/model"
	money "github.com/rockspoon/rs.cor.common-money"
)

// TableBillRequest table bill print request
type TableBillRequest struct {
	Order       order.Order       `json:"order"`
	Check       order.CheckOrder  `json:"check"`
	DineinOrder order.DineInOrder `json:"dineinOrder"`
	Items       []order.OrderItem `json:"items"`
	Language    string            `json:"language"`
}

// ToBill converts the request struct into a printable struct
func (t TableBillRequest) ToBill() Bill {
	bill := Bill{
		RestaurantInfo: t.GetRestaurantInfo(),
		TableInfo:      t.GetTableInfo(),
		InvoiceCheck:   t.GetInvoiceCheck(),
		InvoiceNumber:  t.Order.ID.Hex(),
		BillTime:       t.Order.CreatedAt,
		Items:          t.OrderItemsToInvoiceItems(),
	}
	return bill
}

// GetRestaurantInfo get Restaurant Info from the request
func (t TableBillRequest) GetRestaurantInfo() RestaurantInfo {
	return RestaurantInfo{
		RestaurantName:    t.DineinOrder.Address.Name,
		RestaurantAddress: t.DineinOrder.Address.AddressLine,
		RestaurantZipCode: t.DineinOrder.Address.Zip,
		RestaurantCity:    t.DineinOrder.Address.City,
		RestaurantRegion:  t.DineinOrder.Address.State,
		RestaurantCountry: t.DineinOrder.Address.Country,
		RestaurantPhone:   t.DineinOrder.Address.Phone,
	}
}

// GetTableInfo get TableInfo from the request object
func (t TableBillRequest) GetTableInfo() TableInfo {

	customerName := ""
	if t.DineinOrder.Diners != nil && len(t.DineinOrder.Diners) > 0 {
		customerName = t.DineinOrder.Diners[0].Name
	}
	waiterName := ""
	tableNumber := ""

	if t.DineinOrder.CurrentTableHistory.Tables != nil && len(t.DineinOrder.CurrentTableHistory.Tables) > 0 {
		waiterName = t.DineinOrder.CurrentTableHistory.WaiterName
		tableNumber = t.DineinOrder.CurrentTableHistory.Name
	}

	return TableInfo{
		DiningPartyType: string(t.Order.Type),
		ServerName:      waiterName,
		TableNumber:     tableNumber,
		CustomerName:    customerName,
	}
}

// GetInvoiceCheck get InvoiceCheck from the request
func (t TableBillRequest) GetInvoiceCheck() InvoiceCheck {
	check := t.Check.Summarize(money.CurrencyUSD(), t.DineinOrder.MaxCustomerCount)
	return InvoiceCheck{
		SubTotal:                check.Subtotal,
		DiscountAmount:          check.DiscountAmount,
		DiscountRate:            float32(check.DiscountRate.Value),
		MandatoryGratuityRate:   float32(check.MandatoryGratuityRate.Value),
		MandatoryGratuityAmount: check.MandatoryGratuityAmount,
		TaxRate:                 float32(check.TaxRate.Value),
		TaxAmount:               check.TaxAmount,
		Total:                   check.Total,
	}
}

// OrderItemsToInvoiceItems converts order items into printable models
func (t TableBillRequest) OrderItemsToInvoiceItems() []InvoiceItem {
	items := make([]InvoiceItem, 0)
	for i := 0; i < len(t.Items); i++ {
		name := ""
		for _, n := range t.Items[i].Name {
			if n.Language == t.Language {
				name = n.Value
			}
		}

		mod := ""
		for _, desc := range t.Items[i].Description {
			if desc.Language == t.Language {
				mod = desc.Value
			}
		}
		item := InvoiceItem{
			ItemName: name,
			// TODO review
			Quantity:  1,
			Weight:    1,
			Amount:    t.Items[i].Price,
			Modifiers: mod,
		}
		items = append(items, item)
	}
	return items
}

// Bill printable table bill
type Bill struct {
	RestaurantInfo
	TableInfo
	InvoiceCheck

	InvoiceNumber string
	BillTime      time.Time
	Items         []InvoiceItem
}

// InvoiceCheck check information
type InvoiceCheck struct {
	SubTotal                money.Money
	DiscountAmount          money.Money
	DiscountRate            float32
	MandatoryGratuityRate   float32
	MandatoryGratuityAmount money.Money
	TaxRate                 float32
	TaxAmount               money.Money
	DeliveryFeeAmount       money.Money
	Total                   money.Money
	SalesTaxDescription     string
}

// TableInfo table information
type TableInfo struct {
	DiningPartyType string
	ServerName      string
	TableNumber     string
	CustomerName    string
}

// InvoiceItem ordered item
type InvoiceItem struct {
	ItemName  string
	Quantity  int
	Weight    int
	Amount    money.Money
	Modifiers string
}

// RestaurantInfo restaurant information
type RestaurantInfo struct {
	RestaurantName    string
	RestaurantAddress string
	RestaurantZipCode string
	RestaurantCity    string
	RestaurantRegion  string
	RestaurantCountry string
	RestaurantPhone   string
}
