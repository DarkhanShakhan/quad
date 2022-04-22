package main

import "github.com/01-edu/z01"

func main() {
	h := 5
	v := 5
	for i := 1; i <= h; i++ {
		for j := 1; j <= v; j++ {
			if j == 1 && i == 1 {
				z01.PrintRune('A')
			} else if (j == v && i == 1) || (i == h && j == 1) {
				z01.PrintRune('C')
			} else if j == v && i == h {
				z01.PrintRune('A')
			} else if i == 1 || i == h || j == 1 || j == v {
				z01.PrintRune('B')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
