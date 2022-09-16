package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n > 0 {
		var x []int
		y := 0
		z := 0
		var w int
		for n != 0 {
			y = n % 10
			n /= 10
			x = append(x, y)
		}
		//rentre dans le tableau de rune x[] l'argument n. le nombre revient à étre inversé (8204 -> 4028) mais osef
		z = len(x)
		for i := 0; i < z-1; i++ {
			for j := 0; j < z-i-1; j++ {
				if x[j] > x[j+1] {
					w = x[j]
					x[j] = x[j+1]
					x[j+1] = w
				}
			}
		}
		//arange x[] dans l'ordre avec une variable temporaire w
		for i := 0; i < z; i++ {
			z01.PrintRune(rune(x[i] + 48))
		}
		//imprime x[] avec z = len(s) plus haut
	}
}
