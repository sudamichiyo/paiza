package main

import "fmt"

func main() {
	// 10進数の整数nを標準入力から入力
	var n int
	fmt.Scanf("%d", &n)

	// nを2進数に変換した時の1の個数をカウント
	count := 0
	for n > 0 {
		r := n % 2
		if r == 1 {
			count++
		}
		n = n / 2
	}

	// 1の個数を出力
	fmt.Println(count)
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/count_one_00
