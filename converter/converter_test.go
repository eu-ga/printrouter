package converter

import (
	"testing"

	"github.com/rockspoon/rs.cor.printer-ms/command"
	"github.com/stretchr/testify/require"
)

func TestConverter_Convert(t *testing.T) {
	tt := []struct {
		name        string
		printerType string
		expResult   string
	}{
		{
			name:        "TSP",
			printerType: "TSPP",
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
