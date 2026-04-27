package main

import "fmt"

func main() {
	// メンバーの数n, 仕事の件数mを標準入力から入力
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	// メンバーの名前をn個標準入力から入力
	members := make([]string, 0, n)
	var s string
	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &s)
		members = append(members, s)
	}

	// m件目の仕事を担当するメンバーの名前を出力
	fmt.Println(members[m%n-1])
}

// 問題URL:https://paiza.jp/works/mondai/rank_test_problems_c_0002/rank_test_problems_c_0002__4
