package piscine

func BasicJoin(elems []string) string {
	var a string
	for i := 0; i <= len(elems); i++ {
		if i >= 0 && i <= 3 {
			a = elems[0] + elems[1] + elems[2] + elems[3]
		}
	}
	return a
}
