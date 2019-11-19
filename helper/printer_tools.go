package helper

import (
	"github.com/rockspoon/rs.cor.printer-ms/command"
	"github.com/rockspoon/rs.cor.printer-ms/converter"
)

// GenerateByteCode generate byte array from a list of commands
func GenerateByteCode(commands []command.PrinterCommand, printerType string) []byte {
	if printerType == "TSPP" {
		return converter.TSPPrinterConverter{}.GenerateByteCode(commands)
	}
	return make([]byte, 0)
}
