package kitchen

import (
	"github.com/rockspoon/rs.cor.printer-ms/command"
	"github.com/rockspoon/rs.cor.printer-ms/model"
	"github.com/rockspoon/rs.cor.printer-ms/template/kitchen/tsp"
)

// Generator Kitchen Receipt Generator
type Generator struct{}

// Generate generate receipt command list
func (Generator) Generate(kitchenReceipt model.KitchenReceipt) []command.PrinterCommand {
	return tsp.KitchenReceiptGenerator{}.Generate(kitchenReceipt)
}
