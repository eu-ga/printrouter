package kitchen

import (
	"github.com/rockspoon/rs.cor.printer-ms/command"
	"github.com/rockspoon/rs.cor.printer-ms/model"
	"github.com/rockspoon/rs.cor.printer-ms/template/kitchen/tsp"
)

// Generator Kitchen Receipt Generator
type Generator struct{}

// Generate generate receipt command list according to printer type
func (Generator) Generate(receipt model.KitchenReceipt, printerType string) []command.PrinterCommand {
	var commands []command.PrinterCommand
	if printerType == "TSPP" {
		commands = tsp.ReceiptGenerator{}.Generate(receipt)
	}
	return commands
}
