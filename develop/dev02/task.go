package main

import (
	"strings"
	"unicode"
)

func Unpack(s string) string {
	data := []rune(s)
	if len(data) == 0 || unicode.IsDigit(data[0]) || data[len(data)-1] == '\\' {
		return ""
	}
	var (
		result     strings.Builder
		lastSymbol = data[0]
		lastNumber = -1
	)
	for i := 1; i < len(data); i++ {
		if unicode.IsLetter(data[i]) {
			if lastNumber != -1 {
				for j := 0; j < lastNumber; j++ {
					result.WriteRune(lastSymbol)
				}
				lastNumber = -1
			} else {
				result.WriteRune(lastSymbol)
			}
			lastSymbol = data[i]
		} else if unicode.IsDigit(data[i]) {
			if lastNumber == -1 || lastNumber == 0 {
				lastNumber = int(data[i] - 48)
			} else {
				lastNumber = 10*lastNumber + int(data[i]-48)
			}
		} else if data[i] == '\\' {
			i++
			if !(data[i] == '\\' || unicode.IsDigit(data[i])) {
				return ""
			}
			result.WriteRune(lastSymbol)
			lastSymbol = data[i]
		}
	}
	if lastNumber != -1 {
		for j := 0; j < lastNumber; j++ {
			result.WriteRune(lastSymbol)
		}
	} else if lastSymbol != 0 {
		result.WriteRune(lastSymbol)
	}
	return result.String()
}
