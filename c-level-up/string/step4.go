package main

import "fmt"

func main() {
	// 整数nを標準入力から入力
	var n int
	fmt.Scanf("%d", &n)

	// 整数を3桁で出力(2桁や1桁の場合は０埋めする)
	nstr := fmt.Sprintf("%03d", n)
	fmt.Println(nstr)
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_level_up_problems/c_rank_string_step4
