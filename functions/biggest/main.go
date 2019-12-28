package main

import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
		101, 200, 500,
	}
	fmt.Println(findBiggest(x)) //returns 500
}

func findBiggest(list []int) int {
	var maximum int
	_ = maximum

	for i := 0; i < (len(list) - 1); i++ {
		if list[i] < list[i+1] {
			maximum = list[i+1]
		}
	}

	return maximum
}
