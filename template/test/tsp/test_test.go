package tsp

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rockspoon/rs.cor.common-model/address"
	"github.com/rockspoon/rs.cor.printer-ms/helper"
	"github.com/rockspoon/rs.cor.printer-ms/model"
)

func TestTemplate_Test(t *testing.T) {
	test := model.TestPayload{
		Restaurant: model.RestaurantInfo{
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
		},
	}
	code := helper.GenerateByteCode(TestGenerator{}.Generate(test), "")
	assert.Equal(t, string(code), "")
}
