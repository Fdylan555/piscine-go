package piscine

func Index(s, tofind string) int {
	for i := 0; i <= (len(s) - len(tofind)); i++ {
		if s[i] == tofind[0] {
			count := 0
			for l := 0; l < len(tofind); l++ {
				if tofind[l] == s[i+l] {
					count++
				}
			}
			if count == len(tofind) {
				return i
			}
		}
	}
	return -1
}
