package piscine

func LastRune(s string) rune {
	var last rune
	for _, val := range s {
		last = rune(val)
	}
	return last
}
