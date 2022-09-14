package piscine

func IsUpper(s string) bool {
	bool := true
	for _, val := range s {
		if val >= 'A' && val <= 'Z' {
			bool = true
		} else {
			bool = false
		}
	}
	return bool
}
