package main

import "fmt"

func main() {
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println(x)
	fmt.Println(y)
}

func swap(x *int, y *int) {
	*x, *y = *y, *x
}
