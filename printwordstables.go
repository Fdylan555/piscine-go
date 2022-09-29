package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, i := range a {
		for _, l := range i {
			z01.PrintRune(l)
		}
		z01.PrintRune('\n')
	}
}
