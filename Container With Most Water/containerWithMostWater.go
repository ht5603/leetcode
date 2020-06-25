package main

import (
	"fmt"
)

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	answer := maxArea(height)
	fmt.Println("anser:", answer)
}

func maxArea(height []int) int {
	max := 0
	i := 0
	j := len(height) - 1
	for i < j {
		min := getMin(height[i], height[j])
		max = getMax(max, min*(j-i))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}

func getMin(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

func getMax(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
