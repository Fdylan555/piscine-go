package piscine

func ConcatParams(args []string) string {
	var x string
	for a, i := range args {
		x += i
		if a == len(args)-1 {
			continue
		} else {
			x += "\n"
		}
	}
	return x
}
