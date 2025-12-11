package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int)  {
	i := m - 1
	j := n - 1
	k := m + n - 1
	// Merge from the back
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	// If nums2 has leftovers, copy them
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
	fmt.Println(nums1)
}



func main() {
	merge([]int{1, 2, 0, 0}, 2, []int{5, 6, 7, 8, 9, 10, 11, 12, 13, 11, 1, 1, 1, 1, 1, 1, 1, 1, 1,}, 2)
}