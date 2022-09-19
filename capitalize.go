package piscine

func Capitalize(s string) string {
	x := []rune(s)
	if x[0] >= 97 && x[0] <= 122 {
		x[0] -= 32
	}
	for i := 0; i < len(s); i++ {
		if i != 0 {
			if (x[i] >= 65 && x[i] <= 90) && (x[i-1] >= 32) {
				x[i] += 32
			}
		}
		if (x[i] >= 97 && x[i] <= 122) && ((x[i-1] >= 123) || ((x[i-1] <= 96 && x[i-1] >= 91) || (x[i-1] <= 64 && x[i-1] >= 58) || (x[i-1] <= 47 && x[i-1] >= 32))) {
			x[i] -= 32
		}

	}
	return (string(x))
}
