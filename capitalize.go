package piscine

func Capitalize(s string) string {
	x := []rune(s)
	for i := 0; i < len(s); i++ {
		if x[i] == 32 && x[i+1] >= 97 && x[i+1] <= 122 || x[i] == 43 && x[i+1] >= 97 && x[i+1] <= 122 {
			x[i+1] -= 32
		} else if i == 0 && x[0] >= 97 && x[0] <= 122 {
			x[0] -= 32
		}
	}
	return (string(x))
}
