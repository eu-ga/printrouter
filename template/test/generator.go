package test

import (
	"github.com/rockspoon/rs.cor.printer-ms/command"
	"github.com/rockspoon/rs.cor.printer-ms/model"
	"github.com/rockspoon/rs.cor.printer-ms/template/test/tsp"
)

// Generator Receipt Generator
type Generator struct{}

// Generate generate receipt command list
func (Generator) Generate(test model.TestPayload) []command.PrinterCommand {
	return tsp.TestGenerator{}.Generate(test)
}
