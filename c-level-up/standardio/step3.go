package main

import (
	"fmt"
)

func main() {
	// 整数nを標準入力から入力
	var n int
	fmt.Scanf("%d", &n)

	// n個の整数を標準入力から入力
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	// n個の整数を与えられた順に改行区切りで出力
	for i := range a {
		fmt.Println(a[i])
	}
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_level_up_problems/c_rank_std_in_out_step3
