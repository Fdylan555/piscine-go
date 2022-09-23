package piscine

func StrRev(s string) string {
	var y string
	z := len(s)
	for i := z; i > 0; i-- {
		y += (string(rune(s[i-1])))
	}
	return y
}
