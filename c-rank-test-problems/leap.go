package main

import "fmt"

func main() {
	// 正整数nを標準入力から入力
	var n int
	fmt.Scanf("%d", &n)

	// 西暦年が4で割り切れる年はうるう年
	// ただし、100で割り切れる年はうるう年ではない
	// しかし、400 で割り切れる年はうるう年
	result := ""
	if n%400 == 0 {
		result = "Leap"
	} else if n%100 == 0 {
		result = "Common"
	} else if n%4 == 0 {
		result = "Leap"
	} else {
		result = "Common"
	}

	// Leap（うるう年のとき）または Common（平年のとき）のどちらかを出力
	fmt.Println(result)
}

// 問題URL:https://paiza.jp/works/mondai/rank_test_problems_c_0003/rank_test_problems_c_0003__4
