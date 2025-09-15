package main

import (
	"fmt"
	"strings"
)

func main() {
	// 文字aと文字列sを標準入力から入力
	var a, s string
	fmt.Scanf("%s", &a)
	fmt.Scanf("%s", &s)

	// sにaが含まれているかどうか判定
	// 含まれている場合はYES, 含まれていない場合はNOを出力
	if strings.Contains(s, a) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_level_up_problems/c_rank_string_step2
