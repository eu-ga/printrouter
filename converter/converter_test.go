package converter

import (
	"testing"

	d "github.com/rockspoon/rs.cor.device-model/model"
	"github.com/rockspoon/rs.cor.printer-ms/command"
	"github.com/stretchr/testify/require"
)

func TestConverter_Convert(t *testing.T) {
	tt := []struct {
		name        string
		printerType d.PrinterType
		expResult   string
	}{
		{
			name:        "TSP",
			printerType: d.TSPPrinterType,
			expResult:   "VGhpcyBpcyBhIHRleHQ=",
		},
		{
			name:      "No type",
			expResult: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			converter := NewByteCodeGenerator()
			cmdrs := append(make([]command.PrinterCommand, 0), command.Text("This is a text"))
			res := converter.Convert(cmdrs, tc.printerType)
			require.Equal(t, tc.expResult, res)
		})
	}
}
