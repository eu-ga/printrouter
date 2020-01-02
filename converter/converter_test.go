package converter

import (
	"testing"

	"github.com/rockspoon/rs.cor.printer-ms/command"
	"github.com/stretchr/testify/require"
)

func TestConverter_Convert(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		converter := NewByteCodeGenerator()
		cmdrs := append(make([]command.PrinterCommand, 0), command.Text("This is a text"))
		res := converter.Convert(cmdrs)
		require.Equal(t, "VGhpcyBpcyBhIHRleHQ=", res)
	})
}
