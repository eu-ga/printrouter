package helper

import (
	"strings"
)

func Center(str, padding string, size int) string {
	if len(str) >= size {
		return str
	}
	var builder strings.Builder
	for i := 0; i < (size-len(str))/2; i++ {
		builder.WriteString(padding)
	}
	length, _ := builder.WriteString(str)
	for i := 0; i < (size - length); i++ {
		builder.WriteString(padding)
	}
	return builder.String()
}
