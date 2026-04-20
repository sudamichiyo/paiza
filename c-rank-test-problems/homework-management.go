package main

import "fmt"

func main() {
	// 生徒の人数n, 提出数の基準mを標準入力から入力
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	// i番目の生徒の提出数を標準入力から入力
	submit := make([]int, n)
	for i := 0; i < n; i++ {
		a := 0
		fmt.Scanf("%d", &a)
		submit[i] = a
	}

	// 提出数の内、m以下であるものをカウント
	count := 0
	for _, v := range submit {
		if v <= m {
			count++
		}
	}

	// 結果を出力
	fmt.Println(count)
}

// 問題URL:https://paiza.jp/works/mondai/rank_test_problems_c_0002/rank_test_problems_c_0002__3
