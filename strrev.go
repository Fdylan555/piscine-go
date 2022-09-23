package piscine

func StrRev(s string) string {
	var str string
	z := len(s)
	for i := z; i > 0; i-- {
		str += (string(rune(s[i-1])))
	}
	return s
}
