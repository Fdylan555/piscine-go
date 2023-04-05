package Student

import (
	"math/rand"
)

func Analyze(slot [][]int) int {
	rand.Seed(3)
	a := 0
	for i := 0; i <= len(slot); i++ {
		if len(slot[i]) == len(slot[i+1]) {
			a = 0
		} else {
			a = len(slot[i])
		}
		if slot[i][i] == slot[i+1][i+1] {
			a = 0
		} else {
			a = slot[i][i]
		}
		for j := len(slot); j >= 0; j-- {
			if slot[i][i] == slot[i+1][j-1] {
				a = 0
			} else {
				a = slot[i][j]
			}
		}
	}
	return (a)
}
