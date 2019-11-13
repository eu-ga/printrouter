package model

import (
	money "github.com/rockspoon/rs.cor.common-money"
)

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

// InvoiceItem ordered item
type InvoiceItem struct {
	ItemName  string
	Quantity  int
	Weight    int
	Amount    money.Money
	Modifiers string
}
