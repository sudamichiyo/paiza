package main

import (
	"fmt"
)

func main() {
	// 社員の数である整数nを標準入力から入力
	var n int
	fmt.Scanf("%d", &n)

	// 社員の名前と年齢を標準入力から入力
	// 年齢を＋１する
	name := make([]string, n)
	age := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%s %d", &name[i], &age[i])
		age[i] = age[i] + 1
	}

	// 社員の名前と年齢を出力
	for i := 0; i < n; i++ {
		fmt.Println(name[i], age[i])
	}
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_level_up_problems/c_rank_std_in_out_boss
