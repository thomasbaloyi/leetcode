package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(nums, target))

	nums = []int{3, 2, 4}
	target = 6

	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {

	indices := []int{}

	for i := 0; i < len(nums); i++ {

		for j := i + 1; j < len(nums); j++ {
			if nums[j]+nums[i] == target {
				indices = append(indices, i, j)
				break
			}
		}
	}

	return indices
}
