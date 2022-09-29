package piscine

func Map(f func(int) bool, a []int) []bool {
	var y []bool
	for i := 0; i < len(a); i++ {
		y = append(y, f(a[i]))
	}
	return y
}
