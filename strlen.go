package piscine

func StrLen(s string) int {
	var o int
	for i := 1; i < len(s); i++ {
		if (rune(s[i])) == 96 {
			o = len([]rune(s)) - 1
		} else {
			o = len([]rune(s))
		}
	}
	return o
}
