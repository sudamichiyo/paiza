package main

import (
	"fmt"
)

func main() {
	// 10進数の整数nを標準入力から入力
	var n int
	fmt.Scanf("%d", &n)

	// nを2進数に変換
	b := make([]int, 0, 16)
	for n > 0 {
		b = append(b, n%2)
		n = n / 2
	}

	// sを逆順に並び替え
	s := make([]int, 0, 16)
	for i := 0; i < len(b); i++ {
		s = append(s, b[len(b)-(i+1)])
	}

	// sの桁数をチェック
	if len(s) > 16 {
		fmt.Println("sの桁数が16桁以上です: len(s):", len(s))
		return
	}

	// 1の数を数える
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 1 {
			count++
		}
	}

	// 1の数を出力
	fmt.Println(count)
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/count_one_03
