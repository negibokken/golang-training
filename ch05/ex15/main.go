package main

import (
	"fmt"
)

func main() {
	fmt.Println(max(1, 2, 3, 5))
	fmt.Println(min(1, 2, 3, 5))
	fmt.Println(max2(1, 2, 3, 5))
	fmt.Println(min2(1, 2, 3, 5))
}

func max(nums ...int) (int, error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("arguments length is 0")
	}

	max := nums[0]
	for _, num := range nums {
		if max < num {
			max = num
		}
	}
	return max, nil
}

func min(nums ...int) (int, error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("argmunets length is 0")
	}
	min := nums[0]
	for _, num := range nums {
		if min > num {
			min = num
		}
	}
	return min, nil
}

func max2(n int, nums ...int) int {
	max := n
	for _, num := range nums {
		if max < num {
			max = num
		}
	}
	return max
}

func min2(n int, nums ...int) int {
	min := n
	for _, num := range nums {
		if min > num {
			min = num
		}
	}
	return min
}
