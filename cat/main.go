package main

import (
	"io"
	"io/ioutil"
	"os"

	piscine ".."
	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	var file string
	var chemin string
	var error error
	var data []byte
	var stringError string

	//calcul de len
	var len int
	for range args {
		len++
	}

	//Si pas assez d'arguments
	if len < 2 {
		if _, err := io.Copy(os.Stdin, os.Stdout); err != nil {

		}

		return
	}

	for id := range args {
		if id == 0 {
			continue
		}

		//Lire le fichier
		file = string(args[id])
		chemin = "./" + file
		data, error = ioutil.ReadFile(chemin)

		//Erreur
		if error != nil {
			stringError = "open " + file + ": no such file or directory\n"
			piscine.PrintStr(stringError)
			return
		}

		//Ecrire
		piscine.PrintStr(string(data))
		z01.PrintRune('\n')
	}
}
