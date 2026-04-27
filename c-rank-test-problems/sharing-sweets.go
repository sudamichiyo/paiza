package main

import "fmt"

func main() {
	// お菓子の数nを標準入力から入力
	var n int
	fmt.Scanf("%d", &n)

	// 友達とお菓子を2人でちょうど分け合えるなら、Yes、そうでないなら No を出力
	// つまり、N が偶数なら Yes、奇数なら No と出力してください。
	if n%2 == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

// 問題URL:https://paiza.jp/works/mondai/rank_test_problems_c_0003/rank_test_problems_c_0003__1
