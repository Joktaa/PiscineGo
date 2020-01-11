package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var len int
	tab := os.Args

	for id := range tab {
		id = id
		len++
	}

	for i := len - 1; i > 0; i-- {
		for _, val := range tab[i] {
			z01.PrintRune(val)
		}
		z01.PrintRune('\n')
	}
}
