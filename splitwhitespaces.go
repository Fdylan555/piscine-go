package piscine

func SplitWhiteSpaces(s string) []string {
	var x []string
	var mot string
	for y, i := range s {
		if string(i) == " " {
			if mot != "" {
				x = append(x, mot)
				mot = ""
			}
		} else if y == len(s)-1 {
			mot += string(i)
			x = append(x, mot)
		} else {
			mot += string(i)
		}
	}
	return x
}
