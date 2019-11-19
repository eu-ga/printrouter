package tsp

import (
	"fmt"
	"testing"

	"github.com/rockspoon/rs.cor.printer-ms/helper"
	"github.com/rockspoon/rs.cor.printer-ms/model"
	"github.com/stretchr/testify/require"
)

var TSPTestByteArray = []byte{10, 1, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 10, 10, 74, 97, 110, 32, 49, 44, 32, 48, 48, 48, 49, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 49, 50, 58, 48, 48, 58, 48, 48, 32, 65, 77, 10, 75, 105, 116, 99, 104, 101, 110, 58, 32, 72, 111, 116, 32, 100, 105, 115, 104, 101, 115, 10, 79, 114, 100, 101, 114, 32, 84, 121, 112, 101, 58, 32, 68, 105, 110, 101, 45, 105, 110, 10, 10, 82, 117, 110, 110, 101, 114, 58, 32, 65, 108, 98, 101, 114, 116, 111, 10, 83, 101, 99, 116, 105, 111, 110, 58, 32, 84, 101, 114, 114, 101, 111, 10, 84, 97, 98, 108, 101, 115, 58, 32, 50, 44, 32, 51, 44, 32, 52, 10, 10, 1, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 10, 10, 81, 84, 89, 32, 73, 116, 101, 109, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 70, 105, 114, 101, 32, 84, 121, 112, 101, 10, 10, 32, 49, 120, 32, 67, 104, 101, 115, 115, 66, 117, 114, 103, 101, 114, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 84, 111, 103, 101, 116, 104, 101, 114, 10, 32, 32, 32, 32, 42, 32, 67, 97, 114, 110, 101, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 32, 32, 32, 32, 32, 32, 86, 101, 103, 101, 116, 97, 114, 105, 97, 110, 97, 32, 32, 32, 10, 32, 83, 101, 97, 116, 115, 58, 32, 51, 10, 10, 45, 45, 45, 10, 10, 32, 50, 120, 32, 83, 117, 99, 111, 32, 100, 101, 32, 76, 97, 114, 97, 110, 106, 97, 32, 32, 32, 32, 32, 32, 70, 105, 114, 101, 32, 97, 115, 32, 82, 101, 97, 100, 121, 10, 32, 32, 32, 32, 99, 111, 109, 32, 102, 114, 117, 116, 97, 115, 32, 100, 97, 32, 32, 32, 10, 32, 32, 32, 32, 101, 115, 116, 97, 195, 167, 195, 163, 111, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 32, 83, 101, 97, 116, 115, 58, 32, 49, 44, 32, 53, 10, 10, 45, 45, 45, 10, 10, 32, 49, 120, 32, 72, 97, 109, 98, 117, 114, 103, 101, 114, 32, 115, 117, 112, 101, 114, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 84, 111, 103, 101, 116, 104, 101, 114, 10, 32, 32, 32, 32, 99, 114, 111, 99, 97, 110, 116, 101, 32, 32, 32, 32, 32, 32, 32, 32, 10, 32, 32, 32, 32, 102, 101, 110, 111, 109, 101, 110, 97, 108, 32, 100, 97, 32, 32, 32, 32, 10, 32, 32, 32, 32, 99, 97, 115, 97, 32, 115, 117, 105, 195, 167, 97, 32, 32, 32, 32, 32, 32, 10, 32, 32, 32, 32, 42, 32, 83, 101, 109, 32, 109, 111, 108, 104, 111, 32, 32, 32, 32, 32, 10, 32, 32, 32, 32, 32, 32, 101, 115, 112, 101, 99, 105, 97, 108, 32, 32, 32, 32, 32, 32, 10, 32, 32, 32, 32, 42, 32, 66, 97, 99, 111, 110, 32, 69, 120, 116, 114, 97, 32, 32, 32, 10, 32, 83, 101, 97, 116, 115, 58, 32, 51, 10, 10, 45, 45, 45, 10, 10, 32, 51, 120, 32, 66, 97, 116, 97, 116, 97, 32, 102, 114, 105, 116, 97, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 82, 117, 115, 104, 10, 32, 32, 32, 32, 42, 32, 83, 101, 109, 32, 109, 111, 108, 104, 111, 32, 32, 32, 32, 32, 10, 32, 32, 32, 32, 32, 32, 101, 115, 112, 101, 99, 105, 97, 108, 32, 32, 32, 32, 32, 32, 10, 32, 83, 101, 97, 116, 115, 58, 32, 49, 44, 32, 50, 44, 32, 52, 10, 1, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 10, 10, 10, 10, 10, 10, 10}

func TestTemplate_KitchenTSPReceipt(t *testing.T) {
	code := helper.GenerateByteCode(ReceiptGenerator{}.Generate(getReceipt()), "TSPP")
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

	modifierA := model.KitchenSubEntry{
		Name: "Carne Vegetariana",
	}

	modifierB := model.KitchenSubEntry{
		Name: "Sem molho especial",
	}

	modifierC := model.KitchenSubEntry{
		Name: "Bacon Extra",
	}

	item1 := model.KitchenItem{
		Name:       "ChessBurger",
		Quantity:   1,
		SubEntries: []model.KitchenSubEntry{modifierA},
		FireType:   model.TypesOfFireTogether,
		Seats:      "3",
	}
	item2 := model.KitchenItem{
		Name:       "Suco de Laranja com frutas da estação",
		Quantity:   2,
		SubEntries: []model.KitchenSubEntry{},
		FireType:   model.TypesOfFireAsReady,
		Seats:      "1, 5",
	}

	item3 := model.KitchenItem{
		Name:       "Hamburger super crocante fenomenal da casa suiça",
		Quantity:   1,
		Weight:     1,
		SubEntries: []model.KitchenSubEntry{modifierB, modifierC},
		FireType:   model.TypesOfFireTogether,
		Seats:      "3",
	}

	item4 := model.KitchenItem{
		Name:       "Batata frita",
		Quantity:   3,
		Weight:     1,
		SubEntries: []model.KitchenSubEntry{modifierB},
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
