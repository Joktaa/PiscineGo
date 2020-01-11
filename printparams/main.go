package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	tab := os.Args

	for id := range tab {
		if id == 0 {
			continue
		}
		for _, val := range tab[id] {
			z01.PrintRune(val)
		}
		z01.PrintRune('\n')
	}
}
