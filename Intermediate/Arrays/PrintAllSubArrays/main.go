package main

import "fmt"

func getAllSubArrays(arr []int) [][]int {
	n := len(arr)
	subArrays := make([][]int, n*(n+1)/2)
	idx := 0
	for s := 0; s < n; s++ {
		for e := s; e < n; e++ {
			subArr := make([]int, e-s+1)
			for i := s; i <= e; i++ {
				subArr[i-s] = arr[i]
			}
			fmt.Printf("subarr: %p\n", subArr)
			subArrays[idx] = subArr
			fmt.Printf("subarr[%d]: %p\n", idx, subArr)
			idx++
		}
	}
	return subArrays
}

func main() {
	fmt.Println(getAllSubArrays([]int{2, 8, 9}))
}

// TC:O(N^3)
// SC:O(1); Not using any extra space
