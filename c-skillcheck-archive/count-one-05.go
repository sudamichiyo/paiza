package main

import (
	"fmt"
)

func main() {
	// 10進数の整数nを標準入力から入力
	var n int
	fmt.Scanf("%d", &n)

	// nを2進数にした時各桁が1かどうか判定して、1ならカウント
	count := 0
	for i := 0; i < 16; i++ {
		if testBit(n, i) {
			count++
		}
	}

	// nを2進数にした時の1の数を出力
	fmt.Println(count)
}

// 整数xのn番目のビットが1であるかを判定する関数
func testBit(x, n int) bool {
	return x&(1<<n) != 0
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/count_one_05
