package piscine

func BasicAtoi2(s string) int {
	m := map[rune]int{48: 0, 49: 1, 50: 2, 51: 3, 52: 4, 53: 5, 54: 6, 55: 7, 56: 8, 57: 9}

	var result int = 0
	var valInt int
	var pow int
	var len int

	for i, rune := range s {
		len++
		i = i
		rune = rune
	}

	for i := 0; i < len; i++ {
		if rune(s[i]) < 48 || rune(s[i]) > 57 {
			return 0
		}

		valInt = m[rune(s[i])]

		pow = 1
		for j := 0; j < len-i-1; j++ {
			pow *= 10
		}
		result += valInt * pow
	}

	return result
}
