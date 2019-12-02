package tsp

import (
	"github.com/rockspoon/rs.cor.printer-ms/command"
	"github.com/rockspoon/rs.cor.printer-ms/model"
	"github.com/rockspoon/rs.cor.printer-ms/template"
)

// PaymentReceiptGenerator generates a receipt for tsp printer
type PaymentReceiptGenerator struct {
}

// Generate converts a model into a command list
func (gen PaymentReceiptGenerator) Generate(receipt model.PaymentReceipt) []command.PrinterCommand {
	cmdrs := make([]command.PrinterCommand, 0)

	cmdrs = template.AddRestaurantInfo(receipt.Restaurant, cmdrs)
	cmdrs = template.LineSeparator(cmdrs)
	cmdrs = template.AddServiceInfoBill(receipt.AttendantName, receipt.OrderType, receipt.CreatedAt, receipt.Check, cmdrs)
	cmdrs = template.LineSeparator(cmdrs)
	cmdrs = template.AddItemsBill(receipt.Check.Items, cmdrs)
	cmdrs = template.LineSeparator(cmdrs)
	cmdrs = template.AddCheckTotal(receipt.Check, cmdrs)
	cmdrs = template.LineSeparator(cmdrs)
	cmdrs = template.AddReceiptTotal(receipt.PaymentType, receipt.Paid, receipt.Tips, cmdrs)
	if receipt.Card != nil {
		cmdrs = template.AddReceiptDetails(*receipt.Card, cmdrs)
	}
	cmdrs = template.LineSeparator(cmdrs)
	cmdrs = template.Footer(cmdrs)
	cmdrs = append(cmdrs, command.Cut{})

	return cmdrs
}
