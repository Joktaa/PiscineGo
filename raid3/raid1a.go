package main

import (
	"os"

	piscine ".."
)

func main() {
	x := piscine.Atoi(os.Args[1])
	y := piscine.Atoi(os.Args[2])

	if x == 0 || y == 0 {
		return
	}


	piscine.Raid1a(x, y)
}
