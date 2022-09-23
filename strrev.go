package piscine

func StrRev(s string) string {
	var str string
	z := len(s)
	for i := z - 1; i >= 0; i-- {
		str += (string(rune(s[i])))
	}
	return s
}
