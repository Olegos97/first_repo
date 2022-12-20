package main

import "fmt"

func main() {
	var A int
	var B int
	fmt.Scan(&A, &B)
	sum := 0
	for ; 0 <= A && A <= B && B <= 100; A++ {
		sum += A
	}
	fmt.Println(sum)
}

//camelCase functionNumberTwo
//camel_case function_number_two
