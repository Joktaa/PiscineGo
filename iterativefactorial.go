package piscine

func IterativeFactorial(nb int) int {
	var result int = 1
	if nb >= 0 && nb < 26 {
		for i := 1; i < nb+1; i++ {
			result *= i
		}

		return result
	} else {
		return 0
	}
}
