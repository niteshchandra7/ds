package main

import "fmt"

func getPrefixArray(arr []int, isEven bool) []int {
	prefixSum := make([]int, len(arr))
	start_idx := 0
	if !isEven {
		start_idx = 1
	}
	for i := start_idx; i < len(arr); i++ {
		condition := (isEven && i%2 == 0) || (!isEven && i%2 != 0)
		if condition {
			if i == 0 {
				prefixSum[i] = arr[i]
				continue
			}
			prefixSum[i] = prefixSum[i-1] + arr[i]
		} else {
			prefixSum[i] = prefixSum[i-1]
		}
	}
	return prefixSum
}

func getSpecialIdxCount(arr []int) int {
	PSE := getPrefixArray(arr, true)
	PSO := getPrefixArray(arr, false)
	count := 0
	for idx := 0; idx < len(arr); idx++ {
		SE, SO := 0, 0
		if idx == 0 {
			SE = PSO[len(arr)-1] - PSO[idx]
			SO = PSE[len(arr)-1] - PSE[idx]
		} else {
			SE = PSE[idx-1] + PSO[len(arr)-1] - PSO[idx]
			SO = PSO[idx-1] + PSE[len(arr)-1] - PSE[idx]
		}
		if SO == SE {
			count++
		}
	}
	return count
}

//TC:O(N)
//SC:O(N)

func main() {
	fmt.Println(getSpecialIdxCount([]int{4, 3, 2, 7, 6, -2}))
}
