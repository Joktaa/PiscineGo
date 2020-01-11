package piscine

func StrLen(str string) int {
	var result int
	for i, rune := range str {
		result++
		i = i
		rune = rune
	}
	return result
}
