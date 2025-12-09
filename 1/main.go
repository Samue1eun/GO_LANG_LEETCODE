package main

import (
	"fmt"
)

// nums = [2, 7, 11, 15]
// target = 9
// output = [0, 1]


func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for idx, num := range nums {
		difference := target - num
		val, ok := indexMap[difference]
		if ok {
			return []int{val, idx}
		}
		indexMap[num] = idx
	}
	return []int{}
}



func main() {
	fmt.Println(twoSum([]int{1, 7, 3, 2}, 9))

}