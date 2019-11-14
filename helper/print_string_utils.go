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

// WarpString returns an array os string warped in the lenght warpLength
func WarpString(str string, warpLength int) []string {
	words := strings.Split(str, " ")
	res := []string{""}
	line := 0
	for i := range words {
		if len(words[i])+1+len(res[line]) <= warpLength {
			res[line] = res[line] + words[i] + " "
		} else {
			res[line] = res[line][:(len(res[line]) - 1)] // remove the space form the last word in line
			res = append(res, words[i]+" ")
			line++
		}
	}
	return res
}
