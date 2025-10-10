package main

import (
	"fmt"
	"sort"
)

func main() {
	// 正整数nを標準入力から入力
	var n int
	fmt.Scanf("%d", &n)

	// n個の整数を標準入力から入力
	numbers := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scanf("%d", &num)
		numbers = append(numbers, num)
	}

	// 数を大きい順にソート
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))

	// 結果を改行区切りで出力
	for _, v := range numbers {
		fmt.Println(v)
	}
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_level_up_problems/c_rank_sort_step2
