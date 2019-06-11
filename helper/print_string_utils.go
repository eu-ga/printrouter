package helper

import (
	"strings"
)

// Center returns a string centralized on the line length (size)
func Center(str, padding string, size int) string {
	if len(str) >= size {
		return str
	}
	var builder strings.Builder
	half := (size - len(str)) / 2
	for i := 0; i < half; i++ {
		builder.WriteString(padding)
	}
	length, _ := builder.WriteString(str)
	for i := 0; i < (size - (length + half)); i++ {
		builder.WriteString(padding)
	}
	return builder.String()
}
