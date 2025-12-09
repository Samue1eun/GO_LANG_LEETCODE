package main

import "fmt"


// For this problem what I was to do is create a map and see if there is any duplicates.
// If there are no duplicates, it will return false. If there are two duplicates, it will return true
// In order to find the duplicates, I will need to create a for loop that creates the map.
// Then the if function will need to determine if the there is a duplicate, then print true.
func containsDuplicate(nums []int) bool {
	indexMap := make(map[int]int)
	for _, num := range nums {
		_, ok := indexMap[num]
		if ok {
			if indexMap[num] == 1 {
				return true
			}
		}
		indexMap[num] += 1
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 3, 2, 4}))
}
