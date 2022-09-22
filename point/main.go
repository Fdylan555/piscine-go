package piscine

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func PrintNbr(n int) {
	var tabint []int
	for i := n; i > 0; i /= 10 {
		tabint = append(tabint, i%10)
	}
	for i := len(tabint) - 1; i >= 0; i-- {
		z01.PrintRune(rune(tabint[i] + 48))
	}
}

func main() {
	points := &point{}
	setPoint(points)
	ch := "x = "
	for _, i := range ch {
		z01.PrintRune(i)
	}
	PrintNbr(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	ch = "y = "
	for _, i := range ch {
		z01.PrintRune(i)
	}
	PrintNbr(points.y)
	z01.PrintRune('\n')
}
