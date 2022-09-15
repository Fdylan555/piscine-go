package piscine

func IsPrintable(s string) bool {
	booleen := true
	for _, val := range s {
		if val >= 32 && val <= 127 {
			booleen = true
		} else {
			booleen = false
		}
	}
	return booleen
}
