package piscine

func AlphaCount(s string) int {
	compteur := 0
	val1 := []rune(s)
	for i := range s {
		if val1[i] >= 65 && val1[i] <= 90 || val1[i] >= 97 && val1[i] <= 122 {
			compteur += 1
		}
	}
	return compteur
}
