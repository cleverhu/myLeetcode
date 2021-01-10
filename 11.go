package main

import "fmt"

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	area := 0
	for left < right {
		tmpArea := (right - left) * min(height[right], height[left])
		if tmpArea > area {
			area = tmpArea
		}
		if (height[right] - height[left]) > 0 {
			left++
		} else {
			right--
		}

	}
	return area
}

func main() {
	fmt.Println(maxArea([]int{4, 3, 2, 1, 4}))
}
