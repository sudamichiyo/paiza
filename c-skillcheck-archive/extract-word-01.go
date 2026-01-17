package main

import "fmt"

func main() {
	// 文字列sを標準入力から入力
	var s string
	fmt.Scanf("%s", &s)

	// 文字列sを改行区切りで2回出力
	for i := 0; i < 2; i++ {
		fmt.Println(s)
	}
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/extract_word_01
