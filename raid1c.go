package piscine

import (
	"github.com/01-edu/z01"
)

func Raid1c(x, y int) {
	//VARIABLES
	var nbrX = x - 2
	var nbrY = y - 2
	var i int

	if x < 1 || y < 1 {
		return
	}

	//LIGNE 1
	z01.PrintRune('A')
	if x > 1 {
		for nbrX != 0 {
			z01.PrintRune('B')
			nbrX--
		}
		z01.PrintRune('A')
	}

	//Y
	if y > 1 {
		z01.PrintRune('\n')
		if y > 2 {
			for nbrY != 0 {
				z01.PrintRune('B')
				if x > 1 {
					for i = 0; i < x-2; i++ {
						z01.PrintRune(' ')
					}
					z01.PrintRune('B')
				}
				z01.PrintRune('\n')
				nbrY--
			}
		}

		//LIGNE 2
		nbrX = x - 2
		z01.PrintRune('C')
		if x > 1 {
			for nbrX != 0 {
				z01.PrintRune('B')
				nbrX--
			}
			z01.PrintRune('C')
		}
	}
	z01.PrintRune('\n')
}
