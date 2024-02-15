package main

import "fmt"

func main() {
	nums := []int{4, 3, 5, 6}
	target := 10

	result := TwoSum(nums, target)
	fmt.Printf("%d + %d = %d\n", nums[result[0]], nums[result[1]], target)
	fmt.Println("---------------")
	fmt.Printf("%d index ke %d\n",nums[result[0]],result[0])
	fmt.Printf("%d index ke %d\n",nums[result[1]],result[1])

}

func TwoSum(nums []int, target int) []int {
	MapOfKey:= make(map[int]int)

	for i, num := range nums {
		complement := target - num
		
		if index, ok := MapOfKey[complement]; ok {
			return []int{index, i}
		}
		MapOfKey[num] = i
	}
	return nil
}
