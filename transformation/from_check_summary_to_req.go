package transformation

import (
	// order "github.com/rockspoon/rs.com.order-model/model"
	"strings"

	orderModel "github.com/rockspoon/rs.com.order-model/model"
	summary "github.com/rockspoon/rs.com.order-model/summary"
	money "github.com/rockspoon/rs.cor.common-money"
	model "github.com/rockspoon/rs.cor.printer-ms/transformation/model"
	"golang.org/x/text/language"
)

// FromCheckSummaryToInvoiceCheck transforms a checkSummary into an invoiceCheck
func FromCheckSummaryToInvoiceCheck(checkSummary summary.CheckSummary, language language.Tag) model.InvoiceCheck {
	currency := money.CurrencyUSD()

	feeTitles := []string{}
	taxRate := 0.0
	taxAmount := 0.0
	for i := range checkSummary.EntryTaxies.Entries {
		currency = checkSummary.EntryTaxies.Total.Currency
		if checkSummary.EntryTaxies.Entries[i].AmountOriginal.Type == orderModel.AmountTypePercentage {
			taxRate += checkSummary.EntryTaxies.Entries[i].AmountOriginal.Percentage
		} else {
			taxAmount += checkSummary.EntryTaxies.Entries[i].AmountOriginal.Absolute.Price()
		}
		for j := range checkSummary.EntryTaxies.Entries[i].Title {
			if checkSummary.EntryTaxies.Entries[i].Title[j].Language == language.String() {
				feeTitles = append(feeTitles, checkSummary.EntryTaxies.Entries[i].Title[j].Value)
			}
		}
	}

	discountRate := 0.0
	discountAmount := 0.0
	for i := range checkSummary.EntryDiscounts.Entries {
		currency = checkSummary.EntryDiscounts.Total.Currency
		if checkSummary.EntryDiscounts.Entries[i].AmountOriginal.Type == orderModel.AmountTypePercentage {
			discountRate += checkSummary.EntryDiscounts.Entries[i].AmountOriginal.Percentage
		} else {
			discountAmount += checkSummary.EntryDiscounts.Entries[i].AmountOriginal.Absolute.Price()
		}
	}

	mandatoryGratuityRate := 0.0
	mandatoryGratuityAmount := 0.0
	for i := range checkSummary.EntryServiceCharges.Entries {
		currency = checkSummary.EntryServiceCharges.Total.Currency
		if checkSummary.EntryServiceCharges.Entries[i].AmountOriginal.Type == orderModel.AmountTypePercentage {
			mandatoryGratuityRate += checkSummary.EntryServiceCharges.Entries[i].AmountOriginal.Percentage
		} else {
			mandatoryGratuityAmount += checkSummary.EntryServiceCharges.Entries[i].AmountOriginal.Absolute.Price()
		}
	}

	invoiceCheck := model.InvoiceCheck{
		SubTotal: checkSummary.SubTotal,

		DiscountRate:   float32(discountRate),
		DiscountAmount: money.NewMoney(currency, discountAmount),

		MandatoryGratuityRate:   float32(mandatoryGratuityRate),
		MandatoryGratuityAmount: money.NewMoney(currency, mandatoryGratuityAmount),

		TaxRate:             float32(taxRate),
		TaxAmount:           money.NewMoney(currency, taxAmount),
		SalesTaxDescription: strings.Join(feeTitles, ", "),

		DeliveryFeeAmount: checkSummary.EntryDeliveryFees.Total,

		Total: checkSummary.Total,
	}

	return invoiceCheck
}

func getModifierNames(item summary.EntryItem, language language.Tag) string {
	modifierNames := []string{}
	for i := range item.EntryModifiers.Entries {
		for j := range item.EntryModifiers.Entries[i].Title {
			if item.EntryModifiers.Entries[i].Title[j].Language == language.String() {
				modifierNames = append(modifierNames, item.EntryModifiers.Entries[i].Title[j].Value)
			}
		}
	}
	return strings.Join(modifierNames, ", ")
}

// FromCheckSummaryToInvoiceItems transforms a checkSummary into an array of invoice itens
func FromCheckSummaryToInvoiceItems(checkSummary summary.CheckSummary, language language.Tag) []model.InvoiceItem {
	// currency := money.CurrencyUSD()
	type itemAux struct {
		Quantity  int
		Modifiers string
		Amount    money.Money
	}

	itemMap := make(map[string]itemAux)

	for i := range checkSummary.EntryItems.List {
		itemName := ""
		for j := range checkSummary.EntryItems.List[i].Entry.Title {
			if checkSummary.EntryItems.List[i].Entry.Title[j].Language == language.String() {
				itemName = checkSummary.EntryItems.List[i].Entry.Title[j].Value
			}
		}

		modifiers := getModifierNames(checkSummary.EntryItems.List[i], language)

		item, ok := itemMap[itemName]
		if ok && item.Modifiers == modifiers {
			item.Quantity++
		} else {
			itemMap[itemName] = itemAux{
				Amount:    checkSummary.EntryItems.List[i].GrossPrice,
				Modifiers: modifiers,
				Quantity:  1,
			}
		}
	}

	items := []model.InvoiceItem{}
	for name, value := range itemMap {
		entryItem := model.InvoiceItem{
			ItemName:  name,
			Quantity:  value.Quantity,
			Weight:    1,
			Amount:    value.Amount,
			Modifiers: value.Modifiers,
		}
		items = append(items, entryItem)
	}

	return items
}
