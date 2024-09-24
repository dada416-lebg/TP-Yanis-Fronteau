package main

import (
	"fmt"
)

func main() {
	fmt.Println(FtCoin([]int{1, 2, 5}, 11))
	fmt.Println(FtCoin([]int{2}, 3))
	fmt.Println(FtCoin([]int{1}, 0))

	fmt.Println(Ft_missing([]int{3, 1, 2}))
	fmt.Println(Ft_missing([]int{0, 1}))
	fmt.Println(Ft_missing([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))

	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {2, 3}}))
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {1, 2}, {1, 2}}))

	fmt.Println(Ft_profit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(Ft_profit([]int{7, 6, 4, 3, 1}))
}
