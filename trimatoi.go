package piscine

func TrimAtoi(s string) int {
	w := []rune(s)
	var x int
	var y bool
	var z int
	for i := 0; i < len(s); i++ {
		if w[i] >= 48 && w[i] <= 57 {
			i = len(s)
		} else if w[i] == 45 {
			y = true
		}
	}
	for j := 0; j < len(s); j++ {
		if w[j] >= 48 && w[j] <= 57 {
			x *= 10
			x += (int(w[j] - 48))
		}
	}
	if y == true {
		z = x * 2
		x -= z
	}

	return x
}
