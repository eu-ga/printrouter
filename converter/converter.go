package converter

import (
	"encoding/base64"

	d "github.com/rockspoon/rs.cor.device-model/model"
	"github.com/rockspoon/rs.cor.printer-ms/command"
)

// ByteCodeGenerator printable payload generator
type ByteCodeGenerator struct {
	TSPConverter TSPPrinterConverter
}

// NewByteCodeGenerator creates a new generator
func NewByteCodeGenerator() ByteCodeGenerator {
	return ByteCodeGenerator{TSPConverter: TSPPrinterConverter{}}
}

// Convert convert a list of commands into a base64 encoded string
func (b ByteCodeGenerator) Convert(commands []command.PrinterCommand, printerType d.PrinterType) string {
	var bytes []byte
	switch printerType {
	case d.TSPPrinterType:
		bytes = b.TSPConverter.GenerateByteCode(commands)
	}
	return base64.StdEncoding.EncodeToString(bytes)
}
