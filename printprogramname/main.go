package main

import (
    "os"
    "github.com/01-edu/z01"
)

func main() {
    for i, tt := range os.Args[0] {
        if i > 1 {
            z01.PrintRune(tt)
        }
    }
    z01.PrintRune('\n')
}
