package piscine

func UltimateDivMod(a *int, b *int) {
	result := *a / *b
	remain := *a % *b

	*a = result
	*b = remain
}
