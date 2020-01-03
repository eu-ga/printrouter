package kitchen

import (
	"github.com/rockspoon/rs.cor.printer-ms/command"
	"github.com/rockspoon/rs.cor.printer-ms/model"
	"github.com/rockspoon/rs.cor.printer-ms/template/receipt/tsp"
)

// Generator Receipt Generator
type Generator struct{}

// Generate generate receipt command list
func (Generator) Generate(receipt model.PaymentReceipt) []command.PrinterCommand {
	return tsp.PaymentReceiptGenerator{}.Generate(receipt)
}
