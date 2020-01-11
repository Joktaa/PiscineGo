package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	var tmp int
	var isUp bool
	isFirst := true

	len := 0
	for range tab {
		len++
	}

	for i := 1; i < len; i++ {
		tmp = f(tab[i-1], tab[i])
		if isFirst {
			if tmp == 0 {
				continue
			}
			if tmp > 0 {
				isUp = false
				isFirst = false
				continue
			}
			if tmp < 0 {
				isUp = true
				isFirst = false
				continue
			}
		}

		if tmp > 0 && isUp {
			return false
		}

		if tmp < 0 && !isUp {
			return false
		}
	}
	return true
}
