package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	var a rune = 48
	var b rune = 48
	var c rune = 48
	for a = 48; a < 58; a++ {
		for b = 48; b < 58; b++ {
			for c = 48; c < 58; c++ {
				if a < b && a < c && b < c {
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(c)
					if a != 55 {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
