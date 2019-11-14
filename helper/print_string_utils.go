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

// WarpString Add description
func WarpString(str string, warpLength int) []string {
	originalWords := strings.Split(str, " ")
	formattedWords := []string{}
	for i := range originalWords {
		formattedWords = append(formattedWords, BreakLongWords(originalWords[i], warpLength)...)
	}

	res := []string{""}
	line := 0
	for i := range formattedWords {
		if len(formattedWords[i])+1+len(res[line]) <= warpLength {
			res[line] = res[line] + formattedWords[i] + " "
		} else {
			res[line] = res[line][:(len(res[line]) - 1)] // remove the space form the last word in line
			res = append(res, formattedWords[i]+" ")
			line++
		}
	}
	return res
}

// BreakLongWords Add description
func BreakLongWords(str string, maxChars int) []string {
	res := []string{}
	temp := str
	for len(temp) > maxChars {
		res = append(res, temp[:maxChars-1]+"-")
		temp = temp[maxChars-1:]
	}
	res = append(res, temp)
	return res
}
