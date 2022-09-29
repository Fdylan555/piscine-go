package piscine

func Any(f func(string) bool, a []string) bool {
	var y bool
	for i := 0; i < len(a); i++ {
		if f(a[i]) {
			y = true
		}
	}
	return y
}
