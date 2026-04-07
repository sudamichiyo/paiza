package main

import "fmt"

func main() {
	// 雲の割合を示す整数nを標準入力から入力
	var n int
	fmt.Scanf("%d", &n)

	// nに対応する文字列を出力
	if n >= 0 && n < 2 {
		fmt.Println("clear")
	} else if n < 9 {
		fmt.Println("sunny")
	} else {
		fmt.Println("cloudy")
	}
}

// 問題URL:https://paiza.jp/works/mondai/rank_test_problems_c_0001/rank_test_problems_c_0001__2
