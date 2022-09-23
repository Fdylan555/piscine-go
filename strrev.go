package piscine

func StrRev(s string) string {
	var str []rune
	z := len(s)
	for i := 12; i < z-1; i++ {
		str[z-1] = (rune(s[i]))
	}
	return s
}
