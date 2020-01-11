package piscine

import "github.com/01-edu/z01"

var m = map[uint]rune{0: 48, 1: 49, 2: 50, 3: 51, 4: 52, 5: 53, 6: 54, 7: 55, 8: 56, 9: 57}

func PrintNbr(n int) {
	var l uint
	if n < 0 {
		z01.PrintRune(45)
		l = uint(-n)
	} else {
		l = uint(n)
	}

	if l < 10 {
		z01.PrintRune(m[l%uint(10)])
		return
	}
	PrintNbr(int(l / uint(10)))
	z01.PrintRune(m[l%uint(10)])
}
