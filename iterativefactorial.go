package piscine

func IterativeFactorial(nb int) int { 
	var h int
	a := 1
	if nb >= 40 {
		return 0
	}
	for i := 1; i <= nb; i++ {
		h = int(i)
		a *= h
	}

	return a
}
