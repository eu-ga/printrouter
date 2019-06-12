package tsp

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	d "github.com/rockspoon/rs.cor.device-model/model"
	"github.com/rockspoon/rs.cor.printer-ms/helper"
	"github.com/rockspoon/rs.cor.printer-ms/model"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var TSPTestByteArray = []byte{10, 10, 83, 101, 114, 118, 101, 114, 58, 32, 74, 111, 104, 110, 32, 68, 111, 101, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 83, 116, 97, 116, 105, 111, 110, 58, 32, 83, 116, 97, 116, 105, 111, 110, 32, 35, 49, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 74, 97, 110, 32, 49, 44, 32, 49, 57, 55, 48, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 49, 50, 58, 48, 48, 58, 48, 48, 32, 65, 77, 10, 79, 114, 100, 101, 114, 58, 32, 32, 32, 32, 32, 32, 32, 35, 53, 100, 48, 48, 49, 100, 99, 99, 54, 98, 55, 52, 53, 53, 48, 100, 54, 49, 100, 98, 57, 97, 52, 101, 10, 84, 97, 98, 108, 101, 58, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 49, 50, 51, 52, 53, 54, 10, 65, 100, 100, 114, 101, 115, 115, 58, 32, 84, 101, 115, 116, 32, 97, 100, 100, 114, 101, 115, 115, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 10, 1, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 10, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 65, 115, 32, 82, 101, 97, 100, 121, 10, 81, 84, 89, 32, 32, 73, 116, 101, 109, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 83, 101, 97, 116, 32, 32, 32, 32, 32, 32, 10, 32, 32, 50, 120, 32, 86, 101, 114, 121, 32, 48, 32, 78, 105, 99, 101, 32, 68, 105, 115, 104, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 49, 10, 32, 32, 50, 120, 32, 86, 101, 114, 121, 32, 49, 32, 78, 105, 99, 101, 32, 68, 105, 115, 104, 32, 32, 32, 32, 32, 32, 32, 32, 83, 112, 108, 105, 116, 32, 49, 44, 50, 10, 32, 32, 50, 120, 32, 86, 101, 114, 121, 32, 50, 32, 78, 105, 99, 101, 32, 68, 105, 115, 104, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 32, 32, 32, 32, 69, 120, 116, 114, 97, 32, 78, 105, 99, 101, 10, 10, 1, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 10, 10, 10, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 84, 104, 97, 110, 107, 32, 89, 111, 117, 33, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 80, 111, 119, 101, 114, 101, 100, 32, 98, 121, 32, 82, 111, 99, 107, 115, 112, 111, 111, 110, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 119, 119, 119, 46, 114, 111, 99, 107, 115, 112, 111, 111, 110, 46, 99, 111, 109, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 10, 10, 10, 10, 10}

func TestTemplate_KitchenTSPReceipt(t *testing.T) {
	code := helper.GenerateByteCode(ReceiptGenerator{}.Generate(getReceipt()), d.TSPPrinterType)
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

func getReceipt() model.KitchenReceipt {
	items := make([]model.KitchenItem, 0)
	for i := 0; i < 3; i++ {
		item := model.KitchenItem{
			Name:       "Very " + strconv.Itoa(i) + " Nice Dish",
			Quantity:   2,
			SeatNumber: []string{"1"},
			IsAllSeats: false,
			IsSplit:    false,
		}
		if i == 1 {
			item.IsSplit = true
			item.SeatNumber = []string{"1", "2"}
		}
		if i == 2 {
			item.IsAllSeats = true
			item.Modifiers = "Extra Nice"
		}
		items = append(items, item)
	}

	_id, _ := primitive.ObjectIDFromHex("5d001dcc6b74550d61db9a4e")
	receipt := model.KitchenReceipt{
		Server:             "John Doe",
		Station:            "Station #1",
		FireType:           "As Ready",
		Timestamp:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		InvoiceNumber:      _id.Hex(),
		TableNumber:        "123456",
		IsPrintedForRunner: false,
		OrderType:          "delivery",
		DeliveryAddress:    "Test address",
		Items:              items,
	}
	return receipt
}
