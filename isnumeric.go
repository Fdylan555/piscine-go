package piscine

func IsNumeric(s string) bool {
	ok := true
	for _, val := range s {
		if val < '0' || val > '9' {
			ok = false
		}
	}
	return ok
}
