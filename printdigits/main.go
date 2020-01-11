package main

import (
	"github.com/01-edu/z01"
)

func main() {
	var r rune
	for i := 48; i < 58; i++ {
		r = rune(i)
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
