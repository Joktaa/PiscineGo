package main

import (
	"bufio"
	"fmt"
	"os"
	//"io"
)

func main() {
	x := 0
	y := 0

	str := ""

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		str += scanner.Text()
		
		if x == 0 {
			x = len(scanner.Text())
		}
		y++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
