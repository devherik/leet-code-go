package main

import (
	"fmt"
	"math"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	x, y := len(nums1), len(nums2)
	low, high := 0, x
	for low <= high {
		partitionX := (low + high) / 2
		partitionY := (x+y+1)/2 - partitionX

		maxLeftX := math.MinInt64
		if partitionX > 0 {
			maxLeftX = nums1[partitionX-1]
		}
		// If partitionX is x, nothing on the right side of nums1; use +infinity
		minRightX := math.MaxInt64
		if partitionX < x {
			minRightX = nums1[partitionX]
		}
		// Same for partitionY
		maxLeftY := math.MinInt64
		if partitionY > 0 {
			maxLeftY = nums2[partitionY-1]
		}
		minRightY := math.MaxInt64
		if partitionY < y {
			minRightY = nums2[partitionY]
		}
		// Check if we found the correct partition
		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			// If total length is odd
			if (x+y)%2 != 0 {
				return float64(max(maxLeftX, maxLeftY))
			}
			// If total length is even
			return float64(max(maxLeftX, maxLeftY)+min(minRightX, minRightY)) / 2.0
		} else if maxLeftX > minRightY {
			// We are too far right in nums1, move left
			high = partitionX - 1
		} else {
			// We are too far left in nums1, move right
			low = partitionX + 1
		}
	}
	return 0.0
}

func main04() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
