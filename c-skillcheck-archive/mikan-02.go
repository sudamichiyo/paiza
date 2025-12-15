package main

import "fmt"

func main() {
	// 仕分ける重さの区切りを表す整数n, 与えられる数値の個数を表す整数mを標準入力から入力
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	// m行整数wを標準入力から入力
	numbers := make([]int, 0, m)
	for i := 0; i < m; i++ {
		var w int
		fmt.Scanf("%d", &w)
		numbers = append(numbers, w)
	}

	// 整数wがnの倍数であればYes, そうでなければNoを出力
	for _, w := range numbers {
		if w%n == 0 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/mikan_02
