package converter

import (
	"github.com/rockspoon/rs.cor.printer-ms/command"
)

const (
	dashedLineA = "------------------------------------------------\n"
	dashedLineB = "-------------------------------------\n"
)

// TSPPrinterConverter tsp printer family command converter
type TSPPrinterConverter struct {
}

// GenerateByteCode generates a printable array of bytes from a list of commands
func (TSPPrinterConverter) GenerateByteCode(commands []command.PrinterCommand) []byte {
	result := make([]byte, 0)

	for _, cmdr := range commands {
		switch cmdr.GetType() {
		case command.BitmapCommandType:
			result = append(result, []byte("<BITMAP>"+(cmdr.(command.Bitmap)).Rawb64+"<BITMAP>")...)
		case command.CutCommandType:
			result = append(result, []byte("\n\n\n\n\n")...)
		case command.DashedLineCommandType:
			if (cmdr.(command.Font)).Font == command.FontB {
				result = append(result, []byte(dashedLineB)...)
			} else {
				result = append(result, []byte(dashedLineA)...)
			}
		case command.FontCommandType:
			switch cmdr.(command.Font).Font {
			case command.FontA:
				result = append(result, byte(0))
			case command.FontB:
				result = append(result, byte(1))
			case command.FontC:
				result = append(result, byte(2))
			case command.FontD:
				result = append(result, byte(3))
			case command.FontE:
				result = append(result, byte(4))
			case command.SpecialFontA:
				result = append(result, byte(97))
			case command.SpecialFontB:
				result = append(result, byte(98))
			}
		case command.NewLineCommandType:
			result = append(result, byte(0x0A))
		case command.TextCommandType:
			result = append(result, []byte((cmdr.(command.Text)).Text)...)
		default:
		}
	}
	return result
}
