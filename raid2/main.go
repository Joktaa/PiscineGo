package main

import (
	"fmt"
	"os"
)

func main() {
	params := os.Args[1:]

	if testEntre(params) == false {
		fmt.Println("Error")
		return
	}

	sudoku := remplir2DTab(params)
	fmt.Println(sudoku)
	sudoku = resoudre(sudoku)
	fmt.Println(sudoku)
}

func testEntre(tab []string) bool {
	if len(tab) != 9 {
		return false
	}
	for _, str := range tab {
		if len(str) != 9 {
			return false
		}

		for _, val := range str {
			if val > 57 || val < 49 && val != '.' {
				return false
			}
		}
	}

	return true
}

func resoudre(tab [9][9]string) [9][9]string {
	var aMettre int = 0
	var isOk = true
	var result [9][9]string = tab

	for isNotVide(result) == true {
		for x := 0; x < 9; x++ {
			for y := 0; y < 9; y++ {
				if result[x][y] == "." {
					for nombre := 1; nombre < 10; nombre++ {

						if testLigne(result, x, nombre) == true {
							if testColonne(result, y, nombre) == true {
								if testCarre(result, x, y, nombre) == true {
									aMettre = nombre
									isOk = false
								}
							}
						}
						//mettre le nombre
						if isOk == true && aMettre != 0 {
							fmt.Println(intToString(aMettre))
							result[x][y] = intToString(aMettre)
						}
						isOk = true
						aMettre = 0

					}
				}
			}
		}
	}
	return result
}

func isNotVide(tab [9][9]string) bool {
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if tab[x][y] == "." {
				return true
			}
		}
	}
	return false
}

func remplir2DTab(tab []string) [9][9]string {
	var result [9][9]string
	for x := range tab {
		for y, val := range tab[x] {
			result[x][y] = runeToString(val)
		}
	}
	return result
}

func testLigne(tab [9][9]string, x, nombre int) bool {
	for i := 0; i < 9; i++ {
		if intToString(nombre) == tab[x][i] {
			return false
		}
	}
	return true
}

func testColonne(tab [9][9]string, y, nombre int) bool {
	for i := 0; i < 9; i++ {
		if intToString(nombre) == tab[i][y] {
			return false
		}
	}
	return true
}

func testCarre(tab [9][9]string, posX, posY, nombre int) bool {
	var x int
	var maxX int
	var y int
	var maxY int

	if posX < 3 {
		x = 0
		maxX = 3
	} else if posX > 5 {
		x = 6
		maxX = 9
	} else {
		x = 3
		maxX = 6
	}

	if posY < 3 {
		y = 0
		maxY = 3
	} else if posY > 5 {
		y = 6
		maxY = 9
	} else {
		y = 3
		maxY = 6
	}

	for x < maxX {
		for y < maxY {
			if intToString(nombre) == tab[x][y] {
				return false
			}
			y++
		}
		x++
	}
	return true
}

func intToString(x int) string {
	m := map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9"}
	return m[x]
}

func runeToString(char rune) string {
	m := map[rune]string{'0': "0", '1': "1", '2': "2", '3': "3", '4': "4", '5': "5", '6': "6", '7': "7", '8': "8", '9': "9", '.': "."}
	result := ""
	result += m[char]
	return result
}
