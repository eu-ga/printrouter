package tsp

import (
	"github.com/rockspoon/rs.cor.printer-ms/command"
	"github.com/rockspoon/rs.cor.printer-ms/model"
	"github.com/rockspoon/rs.cor.printer-ms/template"
)

// TestGenerator generates a receipt for tsp printer
type TestGenerator struct {
}

// Generate converts a model into a command list
func (gen TestGenerator) Generate(test model.TestPayload) []command.PrinterCommand {
	cmdrs := make([]command.PrinterCommand, 0)

	cmdrs = template.AddRestaurantInfo(test.Restaurant, cmdrs)
	cmdrs = template.LineSeparator(cmdrs)
	cmdrs = template.Footer(cmdrs)

	return cmdrs
}
