package main

import "fmt"

func getPSArr(arr []int) []int {
	prefixSumArr := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		if i == 0 {
			prefixSumArr[0] = arr[0]
			continue
		}
		prefixSumArr[i] = prefixSumArr[i-1] + arr[i]
	}
	return prefixSumArr
}

func sumOfAllSubArrayUsingPSArray(arr []int) int {
	sum := 0
	prefixSumArr := getPSArr(arr)
	for s := 0; s < len(arr); s++ {
		for e := s; e < len(arr); e++ {
			if s == 0 {
				sum += prefixSumArr[e]
				continue
			}
			sum += (prefixSumArr[e] - prefixSumArr[s-1])
		}
	}
	return sum
	// TC: O(N^2)
	// SC:O(N)
}

func sumOfAllSubArrayUsingContributionTechnique(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += (len(arr) - i) * (i + 1) * arr[i]
	}
	return sum
	// TC: O(N)
	// SC: O(1)
}

func main() {
	fmt.Println(sumOfAllSubArrayUsingPSArray([]int{1, 2, 3}))
	fmt.Println(sumOfAllSubArrayUsingContributionTechnique([]int{1, 2, 3}))

}
