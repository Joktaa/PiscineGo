package piscine

func IsLower(str string) bool {
	for _, val := range str {
		if val < 97 || val > 122 {
			return false
		}
	}
	return true
}
