package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	var a rune = 48
	var b rune = 48
	var c rune = 48
	var d rune = 48
	for a = 48; a < 58; a++ {
		for b = 48; b < 58; b++ {
			for c = 48; c < 58; c++ {
				for d = 48; d < 58; d++ {
					if a*10+b < c*10+d {
						if a*10+b != c*10+d {
							z01.PrintRune(a)
							z01.PrintRune(b)
							z01.PrintRune(' ')
							z01.PrintRune(c)
							z01.PrintRune(d)
							if a == 57 && b == 56 && c == 57 && d == 57 {
								continue
							} else {
								z01.PrintRune(',')
								z01.PrintRune(' ')
							}
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
