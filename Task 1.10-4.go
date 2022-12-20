package main

import "fmt"

func main() {
	var a = 0
	var x = 0
	var i = 0
	for fmt.Scan(&a); a != 0; fmt.Scan(&a) {
		switch {
		case a > x:
			i = 1
			x = a
		case a == x:
			i++
		}
	}
	fmt.Println(i)
}
