package kitchen

import (
	d "github.com/rockspoon/rs.cor.device-model/model"
	"github.com/rockspoon/rs.cor.printer-ms/command"
	"github.com/rockspoon/rs.cor.printer-ms/model"
	"github.com/rockspoon/rs.cor.printer-ms/template/kitchen/tsp"
)

// Generator Kitchen Receipt Generator
type Generator struct{}

// Generate generate receipt command list according to printer type
func (Generator) Generate(receipt model.KitchenReceipt, printerType d.PrinterType) []command.PrinterCommand {
	var commands []command.PrinterCommand
	if printerType == d.TSPPrinterType {
		commands = tsp.ReceiptGenerator{}.Generate(receipt)
	}
	return commands
}
