package tsp

import (
	"github.com/rockspoon/rs.cor.printer-ms/command"
	"github.com/rockspoon/rs.cor.printer-ms/model"
	"github.com/rockspoon/rs.cor.printer-ms/template"
)

// CheckGenerator generates a bill for tsp printer
type CheckGenerator struct {
}

// Generate converts a model into a command list
func (gen CheckGenerator) Generate(bill model.Bill) []command.PrinterCommand {
	cmdrs := make([]command.PrinterCommand, 0)
	for _, check := range bill.Checks {
		cmdrs = template.AddRestaurantInfo(bill.Restaurant, cmdrs)
		cmdrs = template.LineSeparator(cmdrs)

		cmdrs = template.AddServiceInfo(bill.AttendantName, bill.OrderType, check, cmdrs)
		cmdrs = template.LineSeparator(cmdrs)
		cmdrs = template.AddItems(check.Items, cmdrs)
		cmdrs = template.LineSeparator(cmdrs)
		cmdrs = template.AddCheckTotal(check, cmdrs)
		cmdrs = template.LineSeparator(cmdrs)
		cmdrs = template.Footer(cmdrs)
		cmdrs = append(cmdrs, command.Cut{})
	}
	// Header

	return cmdrs
}
