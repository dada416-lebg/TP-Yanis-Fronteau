package main

import "sort"

func FtCoin(coins []int, montant int) int {
	dp := make([]int, montant+1)
	for i := range dp {
		dp[i] = montant + 1
	}
	dp[0] = 0

	for a := 1; a <= montant; a++ {
		for _, piece := range coins {
			if piece <= a {
				if dp[a-piece]+1 < dp[a] {
					dp[a] = dp[a-piece] + 1
				}
			}
		}
	}

	if dp[montant] == montant+1 {
		return -1
	}
	return dp[montant]
}
func Ft_missing(nums []int) int {
	n := len(nums)
	total := n * (n + 1) / 2
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return total - sum
}
func Ft_non_overlap(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	end := intervals[0][1]
	count := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < end {
			count++
		} else {
			end = intervals[i][1]
		}
	}
	return count
}
func Ft_profit(prices []int) int {
	maxProfit := 0
	minPrice := prices[0]
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}
	return maxProfit
}
