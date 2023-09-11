package main

import (
	"fmt"
	"go-algorithm/algorithm/leetcode_3"
	"math"
)

/*
*
121. 买卖股票的最佳时机
*/
func maxProfit(prices []int) int {
	max, min := math.MinInt, math.MaxInt
	for _, v := range prices {
		min = leetcode_3.Min(min, v)
		max = leetcode_3.Max(max, v-min)
	}
	return max
}

//func Min(a, b int) int {
//	if a <= b {
//		return a
//	} else {
//		return b
//	}
//}
//
//func Max(a, b int) int {
//	if a >= b {
//		return a
//	} else {
//		return b
//	}
//}

func main() {
	ans := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(ans))
}
