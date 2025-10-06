package main

import (
	"fmt"
	"strings"
)

func main() {
	// 正整数mを標準入力から入力
	var m int
	fmt.Scanf("%d", &m)

	// m個の文字を標準入力から入力
	chars := make([]string, 0, m)
	for i := 0; i < m; i++ {
		var char string
		fmt.Scanf("%s", &char)
		chars = append(chars, char)
	}

	// 正整数nを標準入力から入力
	var n int
	fmt.Scanf("%d", &n)

	// n個の財産を標準入力から入力
	strs := make([]string, 0, n)
	for i := 0; i < n; i++ {
		var str string
		fmt.Scanf("%s", &str)
		strs = append(strs, str)
	}

	// 各文字が各文字列に対して出現するかを判定
	// 出現する場合はYES, 出現しない場合はNOを出力
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if strings.Contains(strs[j], chars[i]) {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		}
	}
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_level_up_problems/c_rank_for_step4
