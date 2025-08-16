package main

import (
	"fmt"
	"sort"
)

func main() {
	// 整数nを標準入力から入力
	var n int
	fmt.Scanf("%d", &n)

	// n個の数を標準入力から入力
	a := make([]int, n)
	for i := range a {
		fmt.Scanf("%d", &a[i])
	}

	// 配列をソート
	sort.Ints(a)

	// 与えられた数の中で最大値を出力
	fmt.Println(a[len(a)-1])

}

// 問題URL:https://paiza.jp/works/mondai/c_rank_level_up_problems/c_rank_std_in_out_step4
