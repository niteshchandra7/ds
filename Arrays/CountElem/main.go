package main

import "fmt"

func countElem(arr []int, len int) int {
	if len == 0 {
		return len
	}
	maxElem := arr[0]
	count := 0
	for _, elem := range arr {
		if elem > maxElem {
			maxElem = elem
			count = 0
		} else if elem < maxElem {
			continue
		} else {
			count++
		}
	}
	return len - count
}

func main() {
	arr := []int{-3 - 2, 6, 8, 4, 8, 5}
	fmt.Println(countElem(arr, len(arr)))

}

// TC: O(N), SC: O(1)
