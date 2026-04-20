package main

import "fmt"

func main() {
	// 2つの整数a, bを標準入力から入力
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	// aがbより大きければ1を、aがbと等しければ0を、aがbより小さければ-1を格納
	result := 0
	if a > b {
		result = 1
	} else if a < b {
		result = -1
	}

	// 結果を出力
	fmt.Println(result)
}

// 問題URL:https://paiza.jp/works/mondai/rank_test_problems_c_0002/rank_test_problems_c_0002__2
