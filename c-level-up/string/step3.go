package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 文字列sを標準入力から入力
	var s string
	fmt.Scanf("%s", &s)

	// 1番目と4番目を足し合わせたものをa, 2番目と3番目を足し合わせたものをbとする
	a := int(s[0]-'0') + int(s[3]-'0')
	b := int(s[1]-'0') + int(s[2]-'0')

	// 文字列としてのaと文字列としてのbを結合したものを出力
	str := strconv.Itoa(a) + strconv.Itoa(b)
	fmt.Println(str)
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_level_up_problems/c_rank_string_step3
