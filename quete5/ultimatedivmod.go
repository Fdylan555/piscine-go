package piscine

func UltimateDivMod(a *int, b *int) {
	u := *a
	*a = *a / *b
	*b = u % *b
}
