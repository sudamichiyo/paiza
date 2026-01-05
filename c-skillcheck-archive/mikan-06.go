package main

import (
	"fmt"
)

func main() {
	// 仕分ける重さの区切りを表す整数n, みかんの個数を表す整数mを標準入力から入力
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	// みかんの重さを表す整数wを標準入力から入力
	mweight := make([]int, 0, m)
	for i := 0; i < m; i++ {
		var w int
		fmt.Scanf("%d", &w)
		mweight = append(mweight, w)
	}

	// 仕分ける重さが書かれた整数をスライスに格納
	countb := 1000/n + 1
	boxes := make([]int, 0, countb)
	for i := 1; i <= countb; i++ {
		bnum := n * i
		boxes = append(boxes, bnum)
	}

	// 箱に書かれた重さとみかんの重さの絶対値を計算してマップに格納
	b := make(map[int]int, len(boxes))
	for _, mw := range mweight {
		minabs := 1000
		for _, bnum := range boxes {
			abs := Abs(bnum - mw)
			if abs <= minabs {
				minabs = abs
				b[mw] = bnum
			}
		}
	}

	// 結果を出力
	for _, mw := range mweight {
		fmt.Println(b[mw])
	}
}

// 絶対値を求める関数
func Abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/mikan
