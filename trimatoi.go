package piscine

func TrimAtoi(s string) int {
	res := []int{}
	res_final := 0
	len := 0
	len_res := 0
	for i, car := range s {
		len++
		i = i
		car = car
	}
	for val := 0; val < len; val++ {
		if rune(s[val]) >= 48 && rune(s[val]) <= 57 {
			res = append(res, int(s[val]))
			len_res++
		} else if rune(s[val]) == 45 && len_res == 0 {
			res = append(res, 45)
			len_res++
		}
	}
	if len_res == 0 {
		return 0
	}
	for index := 0; index < len_res; index++ {
		if res[index] != 45 {
			res_final += (res[index] - 48) * RecursivePower(10, len_res-(index+1))
		}
	}
	if res[0] == 45 {
		return -res_final
	}
	return res_final
}
