package main

import (
	"fmt"
)

func main() {
	// 10進数の整数nを標準入力から入力
	var n int
	fmt.Scanf("%d", &n)

	// nを2進数にした時最下位の桁が1かどうか判定して、最下位が1ならYes, そうでないならNoを出力
	if n&1 == 1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/count_one_04
