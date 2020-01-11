package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args
	len := 0

	for range args {
		len++
	}

	if len < 2 {
		fmt.Println("File name missing")
		return
	}

	if len > 2 {
		fmt.Println("Too many arguments")
		return
	}

	chemin := "./" + string(args[1])

	data, _ := ioutil.ReadFile(chemin)

	fmt.Println(string(data))
}
