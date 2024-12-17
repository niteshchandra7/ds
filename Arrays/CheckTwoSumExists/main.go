package main

import "fmt"

func CheckTwoSumExists(arr []int, k int) bool {
	mp := make(map[int]interface{})
	for _, elem := range arr {
		_, ok := mp[k-elem]
		if ok {
			return ok
		}
		mp[elem] = nil

	}
	return false
}

func main() {
	fmt.Println(CheckTwoSumExists([]int{3, -2, 1, 4, 3, 6, 8}, 10)) // True
	fmt.Println(CheckTwoSumExists([]int{2, 4, -3, 7}, 8))           // False

}

// TC : O(N)
// SC: O(N)
