package helper

import (
	"github.com/rockspoon/rs.cor.printer-ms/command"
	"github.com/rockspoon/rs.cor.printer-ms/converter"
	"github.com/rockspoon/rs.cor.printer-ms/model"
)

// GenerateByteCode generate byte array from a list of commands
func GenerateByteCode(commands []command.PrinterCommand, printerType model.PrinterType) []byte {
	switch printerType {
	case model.TSPPrinterType:
		return converter.TSPPrinterConverter{}.GenerateByteCode(commands)
	}
	return make([]byte, 0)
}
