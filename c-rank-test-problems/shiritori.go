package main

import "fmt"

func main() {
	// 2つの単語を標準入力から入力
	var s, t string
	fmt.Scanf("%s %s", &s, &t)

	// 1つ目の単語sの末尾と2つ目の単語tの先頭の文字が一致するかチェック
	result := false
	if s[len(s)-1] == t[0] {
		result = true
	}

	// 一致すればYES, そうでなければNOを出力
	if result {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

// 問題URL:https://paiza.jp/works/mondai/rank_test_problems_c_0002/rank_test_problems_c_0002__1
