package tsp

import (
	"fmt"
	"testing"
	"time"

	m "github.com/rockspoon/rs.cor.common-money"
	d "github.com/rockspoon/rs.cor.device-model/model"
	"github.com/rockspoon/rs.cor.printer-ms/helper"
	"github.com/rockspoon/rs.cor.printer-ms/model"
	"github.com/stretchr/testify/require"
)

var TSPTestByteArray = []byte{10, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 82, 101, 115, 116, 97, 117, 114, 97, 110, 116, 32, 78, 97, 109, 101, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 49, 50, 51, 32, 82, 111, 99, 107, 115, 112, 111, 111, 110, 32, 68, 114, 105, 118, 101, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 80, 97, 108, 111, 32, 65, 108, 116, 111, 32, 57, 52, 48, 50, 48, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 67, 65, 32, 85, 83, 65, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 10, 1, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 10, 10, 79, 114, 100, 101, 114, 58, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 35, 49, 50, 51, 52, 53, 54, 55, 54, 56, 10, 10, 74, 97, 110, 32, 49, 44, 32, 49, 57, 55, 48, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 49, 50, 58, 48, 48, 58, 48, 48, 32, 65, 77, 10, 83, 101, 114, 118, 101, 114, 58, 32, 74, 111, 104, 110, 32, 83, 109, 105, 116, 104, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 10, 84, 97, 98, 108, 101, 58, 32, 49, 10, 10, 1, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 10, 10, 81, 84, 89, 32, 32, 73, 116, 101, 109, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 80, 114, 105, 99, 101, 10, 10, 32, 32, 49, 120, 32, 67, 104, 105, 99, 107, 101, 110, 32, 80, 97, 115, 116, 97, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 36, 32, 32, 32, 53, 46, 48, 48, 10, 1, 32, 32, 32, 32, 32, 69, 120, 116, 114, 97, 32, 83, 97, 117, 99, 101, 10, 32, 32, 50, 120, 32, 83, 116, 101, 97, 107, 32, 84, 97, 114, 116, 97, 114, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 36, 32, 32, 32, 53, 46, 48, 48, 10, 1, 32, 32, 32, 32, 32, 76, 111, 111, 107, 44, 32, 106, 117, 115, 116, 32, 98, 101, 99, 97, 117, 115, 101, 32, 73, 32, 100, 111, 110, 39, 116, 32, 98, 101, 32, 103, 10, 32, 32, 32, 32, 32, 105, 118, 105, 110, 103, 32, 110, 111, 32, 109, 97, 110, 32, 97, 32, 102, 111, 111, 116, 32, 109, 97, 115, 115, 97, 103, 101, 32, 100, 111, 110, 10, 32, 32, 32, 32, 32, 39, 116, 32, 109, 97, 107, 101, 32, 105, 116, 32, 114, 105, 103, 104, 116, 32, 102, 111, 114, 32, 77, 97, 114, 115, 101, 108, 108, 117, 115, 32, 10, 32, 32, 32, 32, 32, 116, 111, 32, 116, 104, 114, 111, 119, 32, 65, 110, 116, 119, 111, 110, 101, 32, 105, 110, 116, 111, 32, 97, 32, 103, 108, 97, 115, 115, 32, 104, 10, 32, 32, 32, 32, 32, 111, 117, 115, 101, 10, 10, 1, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 10, 10, 83, 117, 98, 116, 111, 116, 97, 108, 58, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 36, 32, 32, 49, 48, 46, 48, 48, 10, 68, 105, 115, 99, 111, 117, 110, 116, 58, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 36, 32, 32, 45, 49, 46, 48, 48, 10, 71, 114, 97, 116, 117, 105, 116, 121, 32, 40, 49, 48, 46, 48, 48, 41, 58, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 36, 32, 32, 32, 48, 46, 57, 48, 10, 84, 97, 120, 32, 68, 101, 115, 99, 114, 105, 112, 116, 105, 111, 110, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 36, 32, 32, 32, 48, 46, 57, 57, 10, 84, 111, 116, 97, 108, 58, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 36, 32, 32, 32, 57, 46, 57, 48, 10, 10, 1, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 10, 10, 10, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 84, 104, 97, 110, 107, 32, 89, 111, 117, 33, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 80, 111, 119, 101, 114, 101, 100, 32, 98, 121, 32, 82, 111, 99, 107, 115, 112, 111, 111, 110, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 119, 119, 119, 46, 114, 111, 99, 107, 115, 112, 111, 111, 110, 46, 99, 111, 109, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 10, 10, 10, 10, 10}

func TestTemplate_TSPBill(t *testing.T) {
	code := helper.GenerateByteCode(CheckGenerator{}.Generate(getBill()), d.TSPPrinterType)
	t.Log(string(code))
	require.Len(t, code, len(TSPTestByteArray))
	for i := 0; i < len(code); i++ {
		if TSPTestByteArray[i] != code[i] {
			fmt.Print("index ")
			fmt.Print(i)
			fmt.Println(" is different")
		}
		require.Equal(t, TSPTestByteArray[i], code[i])
	}
}

func getBill() model.Bill {
	restaurantInfo := model.RestaurantInfo{
		RestaurantName:    "Restaurant Name",
		RestaurantAddress: "123 Rockspoon Drive",
		RestaurantZipCode: "94020",
		RestaurantCity:    "Palo Alto",
		RestaurantRegion:  "CA",
		RestaurantCountry: "USA",
		RestaurantPhone:   "(555) 444 333",
	}

	tableInfo := model.TableInfo{
		DiningPartyType: "dinein",
		ServerName:      "John Smith",
		TableNumber:     "1",
		CustomerName:    "Jane Doe",
	}

	invoiceCheck := model.InvoiceCheck{
		SubTotal: m.Money{
			Value:    int64(1000),
			Currency: m.CurrencyBRL(),
		},
		DiscountAmount: m.Money{
			Value:    int64(100),
			Currency: m.CurrencyBRL(),
		},
		DiscountRate: float32(0.10),
		MandatoryGratuityAmount: m.Money{
			Value:    int64(90),
			Currency: m.CurrencyBRL(),
		},
		MandatoryGratuityRate: float32(0.10),
		TaxAmount: m.Money{
			Value:    int64(99),
			Currency: m.CurrencyBRL(),
		},
		TaxRate: float32(0.10),
		Total: m.Money{
			Value:    int64(990),
			Currency: m.CurrencyBRL(),
		},
		SalesTaxDescription: "Tax Description",
	}

	items := make([]model.InvoiceItem, 0)

	item1 := model.InvoiceItem{
		ItemName: "Chicken Pasta",
		Quantity: 1,
		Weight:   1,
		Amount: m.Money{
			Value:    int64(500),
			Currency: m.CurrencyBRL(),
		},
		Modifiers: "Extra Sauce",
	}

	item2 := model.InvoiceItem{
		ItemName: "Steak Tartar",
		Quantity: 2,
		Weight:   1,
		Amount: m.Money{
			Value:    int64(250),
			Currency: m.CurrencyBRL(),
		},
		Modifiers: "Look, just because I don't be giving no man a foot massage don't make it right for Marsellus to throw Antwone into a glass house",
	}

	items = append(items, item1, item2)

	bill := model.Bill{
		RestaurantInfo: restaurantInfo,
		TableInfo:      tableInfo,
		InvoiceCheck:   invoiceCheck,
		InvoiceNumber:  "123456768",
		BillTime:       time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		Items:          items,
	}
	return bill
}
