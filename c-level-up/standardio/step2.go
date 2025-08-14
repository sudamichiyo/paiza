package main

import (
	"fmt"
)

func main() {
	// 整数nを標準入力から入力
	var n int
	fmt.Scanf("%d", &n)

	// n回改行区切りでpaizaと出力
	for i := 0; i < n; i++ {
		fmt.Println("paiza")
	}
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_level_up_problems/c_rank_std_in_out_step2
