package main

import "fmt"

func main() {
	// ポイント(自然数A)を標準入力から入力
	var a int
	fmt.Scanf("%d", &a)

	// Aの2倍のポイントを出力
	point := 2 * a
	fmt.Println(point)
}

// 問題URL:https://paiza.jp/works/mondai/rank_test_problems_c_0003/rank_test_problems_c_0003__2
