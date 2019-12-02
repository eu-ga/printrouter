package kitchen

import (
	"github.com/rockspoon/rs.cor.printer-ms/command"
	"github.com/rockspoon/rs.cor.printer-ms/model"
	"github.com/rockspoon/rs.cor.printer-ms/template/receipt/tsp"
)

// Generator Receipt Generator
type Generator struct{}

// Generate generate receipt command list according to printer type
func (Generator) Generate(receipt model.PaymentReceipt, printerType string) []command.PrinterCommand {
	var commands []command.PrinterCommand
	if printerType == "TSPP" {
		commands = tsp.PaymentReceiptGenerator{}.Generate(receipt)
	}
	return commands
}
