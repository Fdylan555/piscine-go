package piscine

func Capitalize(s string) string {
	x := []rune(s)
	for i := 0; i < len(s); i++ {
		if x[i] == 32 && x[i+1] >= 97 && x[i+1] <= 122 || x[i] == 43 && x[i+1] >= 97 && x[i+1] <= 122 {
			x[i+1] -= 32
		} else if i == 0 && x[0] >= 97 && x[0] <= 122 {
			x[0] -= 32
			}else x[i] >= 32 && x[i] <= 47 && x[i] >= 58 && x[i] <= 64 && x[i] >= 91 && x[i] <= 96 && x[i] >= 123 && x[i] <= 127 {
				x[i] >= 65 || x[i] <= 90
			}
		}
	}
	return (string(x))
}
