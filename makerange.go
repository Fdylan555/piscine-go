package piscine

func MakeRange(min, max int) []int {
	var x []int
	cpt := 0
	if min > max {
		return x
	}
	x = make([]int, max-min)
	for i := min; i < max; i++ {
		x[cpt] = i
		cpt += 1
	}
	return x
}
