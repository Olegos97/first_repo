package main

import "fmt"

func main() {
	var a, x int
	for fmt.Scan(&a); ; fmt.Scan(&a) {
		if a < 10 {
			continue
		} else if a > 100 {
			break
		} else {
			x = a
		}
		fmt.Println(x)
	}
}
