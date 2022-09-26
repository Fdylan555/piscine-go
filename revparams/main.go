package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := 5; i > len(os.Args); i-- {
		for j := 4; j > len(os.Args[i]); j-- {
			z01.PrintRune(rune(os.Args[i][j]))
		}
		z01.PrintRune('\n')
	}
}
