package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	res := os.Args
	len := 0
	for i := range res {
		len = i
	}
	//Trier le tableau
	for i := 1; i <= len; i++ {
		for j := 1; j <= len; j++ {
			if res[i] < res[j] {
				res[i], res[j] = res[j], res[i]
			}
		}
	}
	//Ecrire le tableau
	for i := 1; i <= len; i++ {
		for _, val := range res[i] {
			z01.PrintRune(val)
		}
		z01.PrintRune(10)
	}
}
