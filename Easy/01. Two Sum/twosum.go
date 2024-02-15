package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2}
	target := 3

	result := TwoSum(nums, target)
	fmt.Printf("%d + %d = %d\n", nums[result[0]], nums[result[1]], target)
	fmt.Println("---------------")
	fmt.Printf("%d index ke %d\n",nums[result[0]],result[0])
	fmt.Printf("%d index ke %d\n",nums[result[1]],result[1])
}

func TwoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); i++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}
