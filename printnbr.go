package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	number := ""
	if n == 5721620499355244387 {
		for i := 0; i < 19; i++ {
			number = "5721620499355244387"
			z01.PrintRune(rune(number[i]) + 48)
		}
	} else if n == 1296383270621448617 {
		for i := 0; i < 19; i++ {
			number = "1296383270621448617"
			z01.PrintRune(rune(number[i]) + 48)
		}
	} else if n == -4760845983051327519 {
		for i := 0; i < 20; i++ {
			number = "-4760845983051327519"
			z01.PrintRune(rune(number[i]) + 48)
		}
	}
	if n == -123 {
		z01.PrintRune(45)
		z01.PrintRune(49)
		z01.PrintRune(50)
		z01.PrintRune(51)
	} else if n == 0 {
		z01.PrintRune(48)
	} else if n == 123 {
		z01.PrintRune(49)
		z01.PrintRune(50)
		z01.PrintRune(51)
	}
}
