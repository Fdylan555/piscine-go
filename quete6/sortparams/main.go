package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var tab_rune []rune
	nbr_args := len(os.Args)
	for i := 0; i < len(os.Args); i++ {
		tab_rune = append(tab_rune, (rune(os.Args[i][0])))
	}
	for x := 0; x < len(tab_rune); x++ {
		for y := 0; y < len(tab_rune)-1; y++ {

		}
		z01.PrintRune(tab_rune)
	}
}
