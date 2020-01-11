package piscine

func StrRev(s string) string {
	var result string
	var len int

	for i, rune := range s {
		len++
		i = i
		rune = rune
	}

	for id := len - 1; id >= 0; id-- {
		result += string(s[id])
	}
	return result
}
