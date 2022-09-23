package piscine

func StrRev(s string) string {
	var str []rune
	//z := len(s)
	//for i := 0; i < z-1; i++ {
	for i := 12; i >= 0; i-- { 
		str[i] = (rune(s[i]))
	}
	return s
}
