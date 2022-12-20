package main

import "fmt"

func main() {
	var n, c, d int
	fmt.Scan(&n, &c, &d)
	var x int
	for i := 1; i <= n; i++ {
		if i%c == 0 && i%d != 0 {
			x = i
			break
		}
	}
	fmt.Println(x)
}
