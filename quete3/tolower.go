package piscine

func ToLower(s string) string {
	val := []rune(s)
	for i := 0; i < len(val); i++ {
		if val[i] >= 65 && val[i] <= 90 {
			val[i] += 32
		}
	}
	return (string(val))
}
