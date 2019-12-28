package main

import "fmt"

func main() {
	fmt.Println(half(1))
	fmt.Println(half(2))
	fmt.Println(half(3))
}

func half(number int) (int, bool) {
	if number%2 != 0 {
		return number / 2, false
	} else {
		return number / 2, true
	}
}
