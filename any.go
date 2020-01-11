package piscine

func Any(f func(string) bool, arr []string) bool {
	var tmp bool

	for _, val := range arr {
		tmp = f(val)
		if tmp {
			return tmp
		}
	}
	return false
}
