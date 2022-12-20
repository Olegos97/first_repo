package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	sum := 0
	for i := 1; i <= a; i++ {
		fmt.Scan(&b)
		if b%8 == 0 && 10 <= b && b <= 99 {
			sum += b
		}
	}
	fmt.Println(sum)
}
