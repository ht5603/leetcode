package main

// Given nums = [2, 7, 11, 15], target = 9,
// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].
func main() {

}

func twoSum(nums []int, target int) []int {
	result := make(map[int]int)
	answer := make([]int, 2)
	for index, num := range nums {
		_, isExist := result[num]
		if isExist {
			answer[0] = index
			answer[1] = result[num]

		} else {
			result[target-num] = index
		}
	}
	return answer
}
