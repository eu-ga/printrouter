package bill

import (
	"github.com/rockspoon/rs.cor.printer-ms/command"
	"github.com/rockspoon/rs.cor.printer-ms/model"
	"github.com/rockspoon/rs.cor.printer-ms/template/bill/tsp"
)

// Generator Bill Receipt Generator
type Generator struct{}

// Generate generate bill command list
func (Generator) Generate(bill model.Bill) []command.PrinterCommand {
	return tsp.CheckGenerator{}.Generate(bill)
}
