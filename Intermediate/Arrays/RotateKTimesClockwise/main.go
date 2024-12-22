package main

import "fmt"

func rotateKTimesClockwise(arr []int, k int) []int {
	if k <= 1 {
		return arr
	}
	if k > len(arr) {
		k = k % len(arr)
	}
	reverseArray(arr, 0, len(arr)-1)
	reverseArray(arr, k, len(arr)-1)
	reverseArray(arr, 0, k-1)

	return arr
}

func reverseArray(arr []int, s int, e int) []int {
	for s < e {
		temp := arr[s]
		arr[s] = arr[e]
		arr[e] = temp
		s++
		e--
	}
	return arr
}

func main() {
	fmt.Println(rotateKTimesClockwise([]int{3, -2, 1, 4, 6, 9, 8}, 4))
	fmt.Println(rotateKTimesClockwise([]int{4, 1, 6, 9, 2, 14, 7, 8, 3}, 4))
}

// rev(arr): 8 9 6 4 1 -2 3
// rev(arr): 8 7 14 2 9 6 1 4 3
// TC : O(N)
// SC : O(1)
