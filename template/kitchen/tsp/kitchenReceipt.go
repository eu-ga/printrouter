package tsp

import (
	"github.com/rockspoon/rs.cor.printer-ms/command"
	"github.com/rockspoon/rs.cor.printer-ms/model"
	"github.com/rockspoon/rs.cor.printer-ms/template"
)

// KitchenReceiptGenerator generates a receipt for tsp printer
type KitchenReceiptGenerator struct {
}

// Generate converts a model into a command list
func (gen KitchenReceiptGenerator) Generate(kitchenReceipt model.KitchenReceipt) []command.PrinterCommand {
	cmdrs := make([]command.PrinterCommand, 0)
	cmdrs = template.LineSeparator(cmdrs)
	cmdrs = template.AddServiceInfoKitchen(kitchenReceipt, cmdrs)
	cmdrs = template.LineSeparator(cmdrs)
	cmdrs = template.AddItemsKitchen(kitchenReceipt.Items, cmdrs)
	cmdrs = template.LineSeparator(cmdrs)
	cmdrs = append(cmdrs, command.Cut{})

	return cmdrs
}
