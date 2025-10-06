package main

import "fmt"

func main() {
	// 正整数nを標準入力から入力
	var n int
	fmt.Scanf("%d", &n)

	// n個の財産を標準入力から入力
	assets := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var m int
		fmt.Scanf("%d", &m)
		assets = append(assets, m)
	}

	// 金額kを標準入力から入力
	var k int
	fmt.Scanf("%d", &k)

	// 財産がkである人の番号を出力(ただし、そのような人が複数人いる場合にはそれらの人の中で最も小さい番号を出力)
	for i, v := range assets {
		if v == k {
			fmt.Println(i + 1)
			break
		}
	}
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_level_up_problems/c_rank_for_step3
