package main

import (
	"fmt"
)

func main() {
	input := "abcabcbb"
	answer := lengthOfLongestSubstring(input)
	fmt.Println(answer)
}

//      a,b,c,a,b,c,b,b
//index 0,1,2,3,4,5,6,7
//start 0,0,0,1,

func lengthOfLongestSubstring(s string) int {
	start, maxLength := 0, 0
	table := [128]int{}
	for index, _ := range table {
		table[index] = -1
	}
	for i, c := range s {
		fmt.Println("now:", string(c), i)
		if table[c] >= start {
			fmt.Println("flag1, last time index:, start:", table[c], start)
			start = table[c] + 1
			fmt.Println("start變", start)

		}
		table[c] = i
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}

	}

	return maxLength
}
