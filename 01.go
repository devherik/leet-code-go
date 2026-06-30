package main

import "fmt"

func twoSum(nums []int, target int) []int {
	sums := map[int]int{}
	// Populate hash map
	for i, num := range nums {
		sums[num] = i
	}
	// Find complement
	for i, num := range nums {
		if j, ok := sums[target-num]; ok && i != j {
			return []int{i, j}
		}
	}
	return []int{}
}

func main1() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
