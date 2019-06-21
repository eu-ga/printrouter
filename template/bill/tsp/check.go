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
	// Header
	cmdrs = template.AddRestaurantInfo(bill.RestaurantInfo, cmdrs)
	cmdrs = template.LineSeparator(cmdrs)
	cmdrs = template.AddTableInfo(bill.TableInfo, bill.InvoiceNumber, bill.BillTime, cmdrs)

	cmdrs = template.LineSeparator(cmdrs)
	cmdrs = template.AddInvoiceItems(bill.Items, cmdrs)
	cmdrs = template.LineSeparator(cmdrs)
	cmdrs = template.AddInvoiceCheck(bill.InvoiceCheck, cmdrs)
	cmdrs = template.LineSeparator(cmdrs)
	cmdrs = template.Footer(cmdrs)
	cmdrs = append(cmdrs, command.Cut{})

	return cmdrs
}
