package tsp

import (
	"github.com/rockspoon/rs.cor.printer-ms/command"
	"github.com/rockspoon/rs.cor.printer-ms/model"
	"github.com/rockspoon/rs.cor.printer-ms/template"
)

// ReceiptGenerator generates a receipt for tsp printer
type ReceiptGenerator struct {
}

// Generate converts a model into a command list
func (gen ReceiptGenerator) Generate(receipt model.KitchenReceipt) []command.PrinterCommand {
	cmdrs := make([]command.PrinterCommand, 0)
	cmdrs = template.LineSeparator(cmdrs)
	cmdrs = template.AddServiceInfoKitchen(receipt, cmdrs)
	cmdrs = template.LineSeparator(cmdrs)
	cmdrs = template.AddItemsKitchen(receipt.Items, cmdrs)
	cmdrs = template.LineSeparator(cmdrs)
	cmdrs = append(cmdrs, command.Cut{})

	return cmdrs
}
