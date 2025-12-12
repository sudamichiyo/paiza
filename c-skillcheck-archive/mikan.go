package main

import (
	"fmt"
	"math"
)

func main() {
	// 仕分ける重さの区切りを表す整数n, みかんの個数を表す整数mを標準入力から入力
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	// m行みかんの重さを標準入力から入力
	mikanwight := make([]int, 0, m)
	for i := 0; i < m; i++ {
		var mw int
		fmt.Scanf("%d", &mw)
		mikanwight = append(mikanwight, mw)
	}

	// 仕分ける箱をスライスで生成
	boxnum := 1000/n + 1
	boxes := make([]int, boxnum)
	for i := 0; i < boxnum; i++ {
		boxes[i] = n * (i + 1)
	}
	// fmt.Println(boxes)

	// みかんの重さと箱に書かれた重さの二乗差を計算し、差分が最小の箱へ入れる
	// ただし、二乗差の最小値が同じ場合は大きい数字が書かれた箱に入れる
	var inboxindex int
	result := make([]int, m)

	for idx, mw := range mikanwight {
		diffmin := math.Pow(float64(boxes[0]-mikanwight[idx]), 2)
		for i, bw := range boxes {
			diff := math.Pow(float64(bw-mw), 2)
			if diff <= diffmin {
				diffmin = diff
				inboxindex = i
			}
		}
		result[idx] = boxes[inboxindex]
		// fmt.Printf("みかんの重さ: %d, 入れる箱: %d\n", mw, boxes[inboxindex])
	}

	// 結果を出力
	for _, resultbox := range result {
		fmt.Println(resultbox)
	}
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/mikan_00
