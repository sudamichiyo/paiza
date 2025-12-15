package main

import "fmt"

func main() {
	// 仕分ける重さの区切りを表す整数n, みかんの個数を表す整数mを標準入力から入力
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	// m行みかんの重さを標準入力から入力
	mikanweight := make([]int, 0, m)
	for i := 0; i < m; i++ {
		var mw int
		fmt.Scanf("%d", &mw)
		mikanweight = append(mikanweight, mw)
	}

	// n, m, m個のみかんの重さを出力
	fmt.Println(n, m)
	for _, mw := range mikanweight {
		fmt.Println(mw)
	}
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/mikan_01
