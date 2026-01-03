package main

import (
	"fmt"
)

func main() {
	// 仕分ける重さの区切りを表す整数n, みかんの個数を表す整数mを標準入力から入力
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	// みかんの重さを表す整数wを標準入力から入力
	var w int
	fmt.Scanf("%d", &w)

	// 仕分ける重さが書かれた整数をスライスに格納
	var boxes []int
	for i := 1; i <= 1500; i++ {
		if i%n == 0 {
			boxes = append(boxes, i)
		}
	}

	// 箱に書かれた重さとみかんの重さの絶対値を計算してマップに格納
	b := make(map[int]int, len(boxes))
	for _, bnum := range boxes {
		b[bnum] = Abs(bnum - w)
	}

	// 結果を出力
	for _, bnum := range boxes {
		fmt.Println(bnum, b[bnum])
	}
}

// 絶対値を求める関数
func Abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/mikan_05
