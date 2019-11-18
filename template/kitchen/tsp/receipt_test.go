package tsp

import (
	"fmt"
	"testing"

	d "github.com/rockspoon/rs.cor.device-model/model"
	"github.com/rockspoon/rs.cor.printer-ms/helper"
	"github.com/rockspoon/rs.cor.printer-ms/model"
	"github.com/stretchr/testify/require"
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

	tableInfoAlberto := model.DineInKitchenInfo{
		SectionName: "Terreo",
		Tables:      "2, 3, 4",
		RunnerName:  "Alberto",
	}

	modifierA := model.KitchenSubItem{
		Name: "Carne Vegetariana",
	}

	modifierB := model.KitchenSubItem{
		Name: "Sem molho especial",
	}

	modifierC := model.KitchenSubItem{
		Name: "Bacon Extra",
	}

	item1 := model.KitchenItem{
		Name:       "ChessBurger",
		Quantity:   1,
		SubEntries: []model.KitchenSubItem{modifierA},
		FireType:   model.TypesOfFireTogether,
		Seats:      "3",
	}
	item2 := model.KitchenItem{
		Name:       "Suco de Laranja com frutas da estação",
		Quantity:   2,
		SubEntries: []model.KitchenSubItem{},
		FireType:   model.TypesOfFireAsReady,
		Seats:      "1, 5",
	}

	item3 := model.KitchenItem{
		Name:       "Hamburger super crocante fenomenal da casa suiça",
		Quantity:   1,
		Weight:     1,
		SubEntries: []model.KitchenSubItem{modifierB, modifierC},
		FireType:   model.TypesOfFireTogether,
		Seats:      "3",
	}

	item4 := model.KitchenItem{
		Name:       "Batata frita",
		Quantity:   3,
		Weight:     1,
		SubEntries: []model.KitchenSubItem{modifierB},
		FireType:   model.TypesOfFireRush,
		Seats:      "1, 2, 4",
	}

	receipt := model.KitchenReceipt{
		DineInInfo: &tableInfoAlberto,
		OrderType:  model.TypesOfOrderDinein,
		Kitchen:    "Hot dishes",
		Items:      []model.KitchenItem{item1, item2, item3, item4},
	}

	return receipt
}
