package tsp

import (
	"fmt"
	"testing"

	"github.com/rockspoon/rs.cor.common-model/address"
	m "github.com/rockspoon/rs.cor.common-money"
	"github.com/rockspoon/rs.cor.printer-ms/helper"
	"github.com/rockspoon/rs.cor.printer-ms/model"
	"github.com/stretchr/testify/require"
)

var TSPTestByteArray = []byte{10, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 65, 108, 118, 111, 114, 97, 100, 97, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 82, 117, 97, 32, 65, 108, 118, 111, 114, 97, 100, 97, 44, 32, 49, 53, 52, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 83, 195, 163, 111, 32, 80, 97, 117, 108, 111, 32, 48, 53, 53, 57, 51, 48, 49, 48, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 86, 105, 108, 97, 32, 79, 108, 195, 173, 109, 112, 105, 97, 32, 66, 114, 97, 115, 105, 108, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 10, 1, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 10, 10, 74, 97, 110, 32, 49, 44, 32, 48, 48, 48, 49, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 49, 50, 58, 48, 48, 58, 48, 48, 32, 65, 77, 10, 65, 116, 116, 101, 110, 100, 97, 110, 116, 58, 32, 77, 97, 114, 105, 97, 10, 10, 79, 114, 100, 101, 114, 32, 84, 121, 112, 101, 58, 32, 68, 105, 110, 101, 45, 105, 110, 10, 83, 101, 99, 116, 105, 111, 110, 58, 32, 84, 101, 114, 114, 101, 111, 10, 84, 97, 98, 108, 101, 115, 58, 32, 50, 44, 32, 51, 44, 32, 52, 10, 83, 101, 97, 116, 115, 58, 32, 54, 10, 10, 1, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 10, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 85, 110, 105, 116, 121, 32, 32, 32, 70, 105, 110, 97, 108, 10, 81, 84, 89, 32, 73, 116, 101, 109, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 80, 114, 105, 99, 101, 32, 32, 32, 80, 114, 105, 99, 101, 10, 10, 32, 49, 120, 32, 67, 104, 101, 115, 115, 66, 117, 114, 103, 101, 114, 32, 32, 32, 32, 32, 32, 32, 32, 32, 49, 56, 46, 48, 48, 32, 32, 32, 32, 49, 56, 46, 48, 48, 10, 32, 32, 32, 32, 42, 32, 67, 97, 114, 110, 101, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 43, 49, 48, 46, 48, 48, 32, 32, 32, 43, 49, 48, 46, 48, 48, 10, 32, 32, 32, 32, 32, 32, 86, 101, 103, 101, 116, 97, 114, 105, 97, 110, 97, 32, 32, 32, 10, 10, 32, 50, 120, 32, 83, 117, 99, 111, 32, 100, 101, 32, 76, 97, 114, 97, 110, 106, 97, 32, 32, 32, 32, 32, 49, 48, 46, 48, 48, 32, 32, 32, 32, 50, 48, 46, 48, 48, 10, 32, 32, 32, 32, 99, 111, 109, 32, 102, 114, 117, 116, 97, 115, 32, 100, 97, 32, 32, 32, 10, 32, 32, 32, 32, 101, 115, 116, 97, 195, 167, 195, 163, 111, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 10, 1, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 10, 10, 83, 117, 98, 116, 111, 116, 97, 108, 58, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 82, 36, 32, 32, 32, 32, 32, 52, 56, 46, 48, 48, 10, 68, 101, 115, 99, 111, 110, 116, 111, 32, 102, 105, 100, 101, 108, 105, 100, 97, 100, 101, 32, 32, 32, 32, 32, 32, 32, 82, 36, 32, 32, 32, 32, 45, 49, 48, 46, 48, 48, 10, 73, 109, 112, 111, 115, 116, 111, 40, 49, 48, 37, 41, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 82, 36, 32, 32, 32, 32, 32, 32, 51, 46, 56, 48, 10, 84, 111, 116, 97, 108, 58, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 82, 36, 32, 32, 32, 32, 32, 52, 49, 46, 56, 48, 10, 10, 1, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 10, 10, 10, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 84, 104, 97, 110, 107, 32, 89, 111, 117, 33, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 80, 111, 119, 101, 114, 101, 100, 32, 98, 121, 32, 82, 111, 99, 107, 115, 112, 111, 111, 110, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 119, 119, 119, 46, 114, 111, 99, 107, 115, 112, 111, 111, 110, 46, 99, 111, 109, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 10, 10, 10, 10, 10, 10, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 65, 108, 118, 111, 114, 97, 100, 97, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 82, 117, 97, 32, 65, 108, 118, 111, 114, 97, 100, 97, 44, 32, 49, 53, 52, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 83, 195, 163, 111, 32, 80, 97, 117, 108, 111, 32, 48, 53, 53, 57, 51, 48, 49, 48, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 86, 105, 108, 97, 32, 79, 108, 195, 173, 109, 112, 105, 97, 32, 66, 114, 97, 115, 105, 108, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 10, 1, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 10, 10, 74, 97, 110, 32, 49, 44, 32, 48, 48, 48, 49, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 49, 50, 58, 48, 48, 58, 48, 48, 32, 65, 77, 10, 65, 116, 116, 101, 110, 100, 97, 110, 116, 58, 32, 77, 97, 114, 105, 97, 10, 10, 79, 114, 100, 101, 114, 32, 84, 121, 112, 101, 58, 32, 68, 105, 110, 101, 45, 105, 110, 10, 83, 101, 99, 116, 105, 111, 110, 58, 32, 84, 101, 114, 114, 101, 111, 10, 84, 97, 98, 108, 101, 115, 58, 32, 50, 44, 32, 51, 44, 32, 52, 10, 83, 101, 97, 116, 115, 58, 32, 49, 10, 10, 1, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 10, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 85, 110, 105, 116, 121, 32, 32, 32, 70, 105, 110, 97, 108, 10, 81, 84, 89, 32, 73, 116, 101, 109, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 80, 114, 105, 99, 101, 32, 32, 32, 80, 114, 105, 99, 101, 10, 10, 32, 49, 120, 32, 72, 97, 109, 98, 117, 114, 103, 101, 114, 32, 115, 117, 112, 101, 114, 32, 32, 32, 32, 49, 48, 48, 46, 48, 48, 32, 32, 32, 49, 48, 48, 46, 48, 48, 10, 32, 32, 32, 32, 99, 114, 111, 99, 97, 110, 116, 101, 32, 32, 32, 32, 32, 32, 32, 32, 10, 32, 32, 32, 32, 102, 101, 110, 111, 109, 101, 110, 97, 108, 32, 100, 97, 32, 32, 32, 32, 10, 32, 32, 32, 32, 99, 97, 115, 97, 32, 115, 117, 105, 195, 167, 97, 32, 32, 32, 32, 32, 32, 10, 32, 32, 32, 32, 42, 32, 66, 97, 99, 111, 110, 32, 69, 120, 116, 114, 97, 32, 32, 32, 32, 32, 32, 32, 43, 50, 46, 48, 48, 32, 32, 32, 32, 43, 54, 46, 48, 48, 10, 32, 32, 32, 32, 42, 32, 83, 101, 109, 32, 109, 111, 108, 104, 111, 32, 32, 32, 32, 32, 32, 32, 32, 32, 45, 49, 46, 53, 48, 32, 32, 32, 32, 45, 52, 46, 53, 48, 10, 32, 32, 32, 32, 32, 32, 101, 115, 112, 101, 99, 105, 97, 108, 32, 32, 32, 32, 32, 32, 10, 10, 32, 51, 120, 32, 66, 97, 116, 97, 116, 97, 32, 102, 114, 105, 116, 97, 32, 32, 32, 32, 32, 32, 32, 32, 32, 56, 46, 48, 48, 32, 32, 32, 32, 50, 52, 46, 48, 48, 10, 32, 32, 32, 32, 42, 32, 83, 101, 109, 32, 109, 111, 108, 104, 111, 32, 32, 32, 32, 32, 32, 32, 32, 32, 45, 49, 46, 53, 48, 32, 32, 32, 32, 45, 52, 46, 53, 48, 10, 32, 32, 32, 32, 32, 32, 101, 115, 112, 101, 99, 105, 97, 108, 32, 32, 32, 32, 32, 32, 10, 10, 1, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 10, 10, 83, 117, 98, 116, 111, 116, 97, 108, 58, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 82, 36, 32, 32, 32, 32, 49, 50, 49, 46, 48, 48, 10, 73, 109, 112, 111, 115, 116, 111, 40, 49, 48, 37, 41, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 82, 36, 32, 32, 32, 32, 32, 49, 50, 46, 49, 48, 10, 84, 111, 116, 97, 108, 58, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 82, 36, 32, 32, 32, 32, 49, 51, 51, 46, 49, 48, 10, 10, 1, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 10, 10, 10, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 84, 104, 97, 110, 107, 32, 89, 111, 117, 33, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 80, 111, 119, 101, 114, 101, 100, 32, 98, 121, 32, 82, 111, 99, 107, 115, 112, 111, 111, 110, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 119, 119, 119, 46, 114, 111, 99, 107, 115, 112, 111, 111, 110, 46, 99, 111, 109, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 10, 10, 10, 10, 10}

func TestTemplate_TSPBill(t *testing.T) {
	code := helper.GenerateByteCode(CheckGenerator{}.Generate(getBill()), "TSPP")
	t.Log(string(code))
	for i := 0; i < len(code); i++ {
		if TSPTestByteArray[i] != code[i] {
			fmt.Print("index ")
			fmt.Print(i)
			fmt.Println(" is different")
		}
		require.Equal(t, TSPTestByteArray[i], code[i])
	}
}

func moneyInReal(value float64) m.SimpleMoney {
	return m.SimpleMoney{Price: value, Symbol: m.CurrencyBRL().Symbol}
}

func getBill() model.Bill {
	restaurantInfo := model.RestaurantInfo{
		Name: "Alvorada",
		Address: address.Address{
			Name:     "Restaurante Alvorada",
			City:     "São Paulo",
			State:    "SP",
			Country:  "Brasil",
			Address1: "Rua Alvorada, 154",
			Region:   "Vila Olímpia",
			ZipCode:  "05593010",
		},
	}

	tableInfoAlberto := model.DineInOptions{
		SectionName: "Terreo",
		Tables:      "2, 3, 4",
		Seats:       "6",
	}
	tableInfoBernardo := model.DineInOptions{
		SectionName: "Terreo",
		Tables:      "2, 3, 4",
		Seats:       "1",
	}

	modifierA := model.SubEntry{
		Name:       "Carne Vegetariana",
		UnityPrice: moneyInReal(10),
		FinalPrice: moneyInReal(10),
		Index:      3,
	}

	modifierB := model.SubEntry{
		Name:       "Sem molho especial",
		UnityPrice: moneyInReal(-1.5),
		FinalPrice: moneyInReal(-4.5),
		Index:      2,
	}

	modifierC := model.SubEntry{
		Name:       "Bacon Extra",
		UnityPrice: moneyInReal(2),
		FinalPrice: moneyInReal(6),
		Index:      1,
	}

	item1 := model.EntryItem{
		Name:       "ChessBurger",
		Quantity:   1,
		UnityPrice: moneyInReal(18),
		FinalPrice: moneyInReal(18),
		SubEntries: []model.SubEntry{modifierA},
	}
	item2 := model.EntryItem{
		Name:       "Suco de Laranja com frutas da estação",
		Quantity:   2,
		UnityPrice: moneyInReal(10),
		FinalPrice: moneyInReal(20),
		SubEntries: []model.SubEntry{},
	}

	item3 := model.EntryItem{
		Name:       "Hamburger super crocante fenomenal da casa suiça",
		Quantity:   1,
		Weight:     1,
		UnityPrice: moneyInReal(100),
		FinalPrice: moneyInReal(100),
		SubEntries: []model.SubEntry{modifierB, modifierC},
	}

	item4 := model.EntryItem{
		Name:       "Batata frita",
		Quantity:   3,
		Weight:     1,
		UnityPrice: moneyInReal(8),
		FinalPrice: moneyInReal(24),
		SubEntries: []model.SubEntry{modifierB},
	}

	charge1 := model.SubEntry{
		Name:        "Imposto",
		Description: "10%",
		Index:       3,
		FinalPrice:  moneyInReal(3.8),
	}

	charge2 := model.SubEntry{
		Name:       "Desconto fidelidade",
		Index:      2,
		FinalPrice: moneyInReal(-10),
	}

	charge3 := model.SubEntry{
		Name:        "Imposto",
		Description: "10%",
		Index:       1,
		FinalPrice:  moneyInReal(12.1),
	}

	check1 := model.Check{
		DineInOptions: &tableInfoAlberto,
		Items:         []model.EntryItem{item1, item2},
		Subtotal:      moneyInReal(48),
		Charges:       []model.SubEntry{charge1, charge2},
		Total:         moneyInReal(41.8),
	}

	check2 := model.Check{
		DineInOptions: &tableInfoBernardo,
		Items:         []model.EntryItem{item3, item4},
		Subtotal:      moneyInReal(121),
		Charges:       []model.SubEntry{charge3},
		Total:         moneyInReal(133.1),
	}

	bill := model.Bill{
		Restaurant:    restaurantInfo,
		OrderType:     model.TypesOfOrderDinein,
		AttendantName: "Maria",
		Checks:        []model.Check{check1, check2},
	}
	return bill
}
