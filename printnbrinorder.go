package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var m = map[rune]int{0: 48, 1: 49, 2: 50, 3: 51, 4: 52, 5: 53, 6: 54, 7: 55, 8: 56, 9: 57}
	res := []int{}
	len := 0
	//Créer le tableau de i
	if n <= 0 {
		res = append(res, 0)
	} else if n > 0 {
		for n > 0 {
			temp := n % 10
			res = append(res, temp)
			n /= 10
		}
	}
	//Trouver la longueur du tableau
	for index, val := range res {
		len++
		index = index
		val = val
	}
	//Trier le tableau
	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {
			if res[i] < res[j] {
				tampon := res[j]
				res[j] = res[i]
				res[i] = tampon
			}
		}
	}
	//PrintRune la liste
	for p := 0; p < len; p++ {
		z01.PrintRune(rune(m[rune(res[p])]))
	}
}
