package piscine

func IsUpper(s string) bool {
	ok := true
	for _, val := range s {
		if val <= 'A' || val >= 'Z' {
			ok = false
		}
	}
	return ok
}
