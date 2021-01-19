package main

import (
	"fmt"
	"strconv"
)

func Series(n int) int {
	var f []int
	f = make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	var i int
	for i = 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

func FibonacciNumber(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciNumber(n-1) + FibonacciNumber(n-2)
}

func main() {
	var i int
	for i = 0; i <= 9; i++ {
		fmt.Print(strconv.Itoa(Series(i)) + " ")
	}
	fmt.Println("")
	for i = 0; i <= 9; i++ {
		fmt.Print(strconv.Itoa(FibonacciNumber(i)) + " ")
	}
	fmt.Println("")
}
