package main

import "fmt"

func main() {
	// 手荷物の個数n, 重さの基準値mを標準入力から入力
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	// 荷物の各重量を標準入力から入力
	var bw int
	bagweights := make([]int, 0, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &bw)
		bagweights = append(bagweights, bw)
	}

	// 荷物の総重量を計算
	total := 0
	for _, v := range bagweights {
		total += v
	}

	// 基準値と比較して結果を出力
	if total <= m {
		fmt.Println("OK")
	} else {
		fmt.Println("NG")
	}

}

// 問題URL:https://paiza.jp/works/mondai/rank_test_problems_c_0001/rank_test_problems_c_0001__3
