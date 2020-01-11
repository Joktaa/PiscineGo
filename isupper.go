package piscine

func IsUpper(str string) bool {
	for _, val := range str {
		if val < 65 || val > 90 {
			return false
		}
	}
	return true
}
