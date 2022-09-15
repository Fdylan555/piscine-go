package piscine

import "github.com/01-edu/z01"

func ToLower(s string) string {
	for _, val := range s {
		z01.PrintRune(rune(val))
	}
	return ""
}
