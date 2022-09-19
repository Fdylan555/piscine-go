package piscine

func IterativeFactorial(nb int) int {
	var a int
	var h int
	a = 0
	h = 1
	for i := 1; i <= nb; i++ {
		 h *= i 

	a = h 
	}

	return a
}
