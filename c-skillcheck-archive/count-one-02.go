package main

import (
	"fmt"
)

func main() {
	// 2進数の整数sを標準入力から入力
	var s string
	fmt.Scanf("%s", &s)

	// inputの中身を文字列に変換
	if len(s) > 16 {
		fmt.Println("sの桁数が16桁以上です: len(s):", len(s))
		return
	}

	// 1の数を数える
	count := 0
	for i := 0; i < len(s); i++ {
		if (s[i] - 48) == 1 {
			count++
		}
	}

	// 1の数を出力
	fmt.Println(count)
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/count_one_02
