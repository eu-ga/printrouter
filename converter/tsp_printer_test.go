package converter

import (
	"testing"

	"github.com/rockspoon/rs.cor.printer-ms/command"
	"github.com/stretchr/testify/require"
)

func TestTSPPrinter_Bitmap(t *testing.T) {
	tt := []struct {
		name      string
		data      []command.PrinterCommand
		expResult []byte
	}{
		{
			name:      "BitmapCommandType",
			data:      []command.PrinterCommand{command.Bitmap{Rawb64: "ASD"}},
			expResult: []byte("<BITMAP>ASD<BITMAP>"),
		},
		{
			name:      "CutCommandType",
			data:      []command.PrinterCommand{command.Cut{}},
			expResult: []byte("\n\n\n\n\n"),
		},
		{
			name:      "BitmapCommandType, CutCommandType",
			data:      []command.PrinterCommand{command.Bitmap{Rawb64: "ASD"}, command.Cut{}},
			expResult: []byte("<BITMAP>ASD<BITMAP>\n\n\n\n\n"),
		},
		{
			name:      "DashedLineCommandType: command.FontB",
			data:      []command.PrinterCommand{command.DashedLine{Font: command.FontB}},
			expResult: []byte("-------------------------------------\n"),
		},
		{
			name:      "DashedLineCommandType: command.FontA",
			data:      []command.PrinterCommand{command.DashedLine{Font: command.SpecialFontA}},
			expResult: []byte("------------------------------------------------\n"),
		},
		{
			name:      "FontCommandType: FontA",
			data:      []command.PrinterCommand{command.FontA},
			expResult: []byte{byte(0)},
		},
		{
			name:      "FontCommandType: FontB",
			data:      []command.PrinterCommand{command.FontB},
			expResult: []byte{byte(1)},
		},
		{
			name:      "FontCommandType: FontC",
			data:      []command.PrinterCommand{command.FontC},
			expResult: []byte{byte(2)},
		},
		{
			name:      "FontCommandType: FontD",
			data:      []command.PrinterCommand{command.FontD},
			expResult: []byte{byte(3)},
		},
		{
			name:      "FontCommandType: FontE",
			data:      []command.PrinterCommand{command.FontE},
			expResult: []byte{byte(4)},
		},
		{
			name:      "FontCommandType: SpecialFontA",
			data:      []command.PrinterCommand{command.SpecialFontA},
			expResult: []byte{byte(97)},
		},
		{
			name:      "FontCommandType: SpecialFontB",
			data:      []command.PrinterCommand{command.SpecialFontB},
			expResult: []byte{byte(98)},
		},
		{
			name:      "NewLineCommandType",
			data:      []command.PrinterCommand{command.NewLine{}},
			expResult: []byte{byte(0x0A)},
		},
		{
			name:      "TextCommandType",
			data:      []command.PrinterCommand{command.Text("This is a text")},
			expResult: []byte("This is a text"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			converter := TSPPrinterConverter{}
			res := converter.GenerateByteCode(tc.data)
			require.Equal(t, tc.expResult, res)
		})
	}
}
