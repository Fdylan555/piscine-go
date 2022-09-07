package main

import "github.com/01-edu/z01"

func main() {
	for compteur := 48; compteur <= 57; compteur++ {
		z01.PrintRune(rune(compteur))
	}
	z01.PrintRune('\n')
}
