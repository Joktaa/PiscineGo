package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(str string) {
	for i, rune := range str {
		z01.PrintRune(rune)
		i = i
	}
}
