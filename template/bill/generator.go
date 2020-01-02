package bill

import (
	"github.com/rockspoon/rs.cor.printer-ms/command"
	"github.com/rockspoon/rs.cor.printer-ms/model"
	"github.com/rockspoon/rs.cor.printer-ms/template/bill/tsp"
)

// Generator Bill Receipt Generator
type Generator struct{}

// Generate generate bill command list according to printer type
func (Generator) Generate(bill model.Bill, printerType string) []command.PrinterCommand {
	var commands []command.PrinterCommand
	if printerType == "TSP" {
		commands = tsp.CheckGenerator{}.Generate(bill)
	}
	return commands
}
