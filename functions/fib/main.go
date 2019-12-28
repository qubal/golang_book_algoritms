package main

import "fmt"

func main() {
	fmt.Println(fib(3))
}

func fib(n int) int {
	if n < 2 {
		return n
	} else {
		return fib(n-1) + fib(n-2)
	}
}
