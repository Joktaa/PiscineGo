package piscine

func Atoi(s string) int {
	m := map[rune]int{43: 11, 45: 10, 48: 0, 49: 1, 50: 2, 51: 3, 52: 4, 53: 5, 54: 6, 55: 7, 56: 8, 57: 9}

	var result int = 0
	var valInt int
	var pow int
	var len int
	isSigned := 0

	for i, rune := range s {
		len++
		i = i
		rune = rune
	}

	for i := 0; i < len; i++ {
		if rune(s[i]) == 43 || rune(s[i]) == 45 {
			if i != 0 {
				return 0
			}
			isSigned = m[rune(s[i])]
		} else if rune(s[i]) < 48 || rune(s[i]) > 57 {
			return 0
		} else {
			valInt = m[rune(s[i])]
		}

		pow = 1
		for j := 0; j < len-i-1; j++ {
			pow *= 10
		}
		result += valInt * pow
	}
	if isSigned == 10 {
		return -result
	} else {
		return result
	}
}
