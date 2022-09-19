package piscine

func IterativePower(nb int, power int) int {
	t := 1
	if nb >= 21 {
		return 0
	}
	if power < 0 {
		return 0
	}	
	for power != 0 {
		t *= nb
		power -= 1
	}
	return t
}
