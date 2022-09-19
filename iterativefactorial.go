package piscine

func IterativeFactorial(nb int) int {
	var a int
	a = 1
	for i := 1; i <= nb; i++ {
		a *= i
	}
	return a
}
