package helper

import (
	d "github.com/rockspoon/rs.cor.device-model/model"
	"github.com/rockspoon/rs.cor.printer-ms/command"
	"github.com/rockspoon/rs.cor.printer-ms/converter"
)

// GenerateByteCode generate byte array from a list of commands
func GenerateByteCode(commands []command.PrinterCommand, printerType d.PrinterType) []byte {
	if printerType == d.TSPPrinterType {
		return converter.TSPPrinterConverter{}.GenerateByteCode(commands)
	}
	return make([]byte, 0)
}
