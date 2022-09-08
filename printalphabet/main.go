package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for compteur := 97; compteur <= 122; compteur++ {
		z01.PrintRune(rune(compteur))
	}
	z01.PrintRune('\n')
}

//for i:= (a=1; i z=26; augmenter +1)
//z01.PrintRune(i)
