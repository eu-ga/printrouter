package model

import (
	"time"

	order "github.com/rockspoon/rs.com.order-model/model"
	summary "github.com/rockspoon/rs.com.order-model/summary"
	transform "github.com/rockspoon/rs.cor.printer-ms/transformation"
	model "github.com/rockspoon/rs.cor.printer-ms/transformation/model"
	vm "github.com/rockspoon/rs.cor.venue-model/v4/model"
	"golang.org/x/text/language"
)

// TableBillRequest table bill print request
type TableBillRequest struct {
	Order    order.Order          `json:"order"`
	Check    summary.CheckSummary `json:"check"`
	Venue    vm.Address           `json:"venue"`
	Language language.Tag         `json:"language"`
}

// ToBill converts the request struct into a printable struct
func (t TableBillRequest) ToBill() Bill {
	bill := Bill{
		RestaurantInfo: t.GetRestaurantInfo(),
		TableInfo:      t.GetTableInfo(),
		InvoiceCheck:   transform.FromCheckSummaryToInvoiceCheck(t.Check, t.Language),
		//InvoiceNumber:  t.Order.ID.Hex(),
		BillTime: t.Order.CreatedAt,
		Items:    transform.FromCheckSummaryToInvoiceItems(t.Check, t.Language), // t.OrderItemsToInvoiceItems(),
	}
	return bill
}

// GetRestaurantInfo get Restaurant Info from the request
func (t TableBillRequest) GetRestaurantInfo() RestaurantInfo {
	return RestaurantInfo{
		RestaurantName:    t.Venue.Name,     //t.DineinOrder.Address.Name,
		RestaurantAddress: t.Venue.Address1, //t.DineinOrder.Address.AddressLine,
		RestaurantZipCode: t.Venue.ZipCode,  //t.DineinOrder.Address.Zip,
		RestaurantCity:    t.Venue.City,     //t.DineinOrder.Address.City,
		RestaurantRegion:  t.Venue.State,    //t.DineinOrder.Address.State,
		RestaurantCountry: t.Venue.Country,  //t.DineinOrder.Address.Country,
		RestaurantPhone:   "",               //t.DineinOrder.Address.Phone,
	}
}

// GetTableInfo get TableInfo from the request object
func (t TableBillRequest) GetTableInfo() TableInfo {

	customerName := ""
	if t.Check.CheckPayer != nil {
		customerName = t.Check.CheckPayer.Name
	}

	waiterName := ""
	elementNames := []string{}
	if t.Order.DineInOrder != nil {
		waiterName = t.Order.DineInOrder.AssignedTo.EmployeeID.Hex()
		for i := range t.Order.DineInOrder.FloorPlanLocation.FloorPlanElements {
			elementNames = append(elementNames, t.Order.DineInOrder.FloorPlanLocation.FloorPlanElements[i].FloorPlanElementName)
		}
	}

	return TableInfo{
		DiningPartyType: string(t.Order.Type),
		ServerName:      waiterName,
		TableNumber:     elementNames,
		CustomerName:    customerName,
	}
}

// GetInvoiceCheck get InvoiceCheck from the request
// func (t TableBillRequest) GetInvoiceCheck() model.InvoiceCheck {
// 	check := t.Check.Summarize(money.CurrencyUSD(), t.DineinOrder.MaxCustomerCount)
// 	return model.InvoiceCheck{
// 		SubTotal:                check.Subtotal,
// 		DiscountAmount:          check.DiscountAmount,
// 		DiscountRate:            float32(check.DiscountRate.Value),
// 		MandatoryGratuityRate:   float32(check.MandatoryGratuityRate.Value),
// 		MandatoryGratuityAmount: check.MandatoryGratuityAmount,
// 		TaxRate:                 float32(check.TaxRate.Value),
// 		TaxAmount:               check.TaxAmount,
// 		Total:                   check.Total,
// 	}
// }

// OrderItemsToInvoiceItems converts order items into printable models
// func (t TableBillRequest) OrderItemsToInvoiceItems() []model.InvoiceItem {
// 	items := make([]model.InvoiceItem, 0)
// 	for i := 0; i < len(t.Items); i++ {
// 		name := ""
// 		for _, n := range t.Items[i].Name {
// 			if n.Language == t.Language {
// 				name = n.Value
// 			}
// 		}

// 		mod := ""
// 		for _, desc := range t.Items[i].Description {
// 			if desc.Language == t.Language {
// 				mod = desc.Value
// 			}
// 		}
// 		item := model.InvoiceItem{
// 			ItemName: name,
// 			// TODO review
// 			Quantity:  1,
// 			Weight:    1,
// 			Amount:    t.Items[i].Price.Amount,
// 			Modifiers: mod,
// 		}
// 		items = append(items, item)
// 	}
// 	return items
// }

// Bill printable table bill
type Bill struct {
	RestaurantInfo
	TableInfo
	InvoiceCheck model.InvoiceCheck

	InvoiceNumber string
	BillTime      time.Time
	Items         []model.InvoiceItem
}

// TableInfo table information
type TableInfo struct {
	DiningPartyType string
	ServerName      string
	TableNumber     []string
	CustomerName    string
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
