package main

import (
	"fmt"
	"sort"
)

func main() {
	intList := []int{0, 0, 0}
	answer := threeSum(intList)
	fmt.Println(answer)
}

func threeSum(nums []int) [][]int {
	answer := make([][]int, 0)
	n := len(nums)
	if n == 0 || n < 3 {
		return answer
	}
	sort.Ints(nums) //先排序
	//[-4 -1 -1 0 1 2]
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		low := i + 1
		high := len(nums) - 1
		sum := -nums[i]
		for low < high {
			if sum == nums[low]+nums[high] {
				answer = append(answer, []int{nums[low], nums[high], nums[i]})

				for nums[low] == nums[low+1] {
					low++
				}
				for nums[high] == nums[high-1] {
					high--
				}
				low++
				high--
			} else if sum < nums[low]+nums[high] {
				high--
			} else {
				low++
			}
		}
	}
	return answer
}
