package main

import (
	"fmt"
)

func main() {
	// 生鮮食品の量m[kg], p[%], q[%]を標準入力から入力
	var m, p, q int
	fmt.Scanf("%d %d %d", &m, &p, &q)

	// 生で売れた分を計算
	m = m * 1000000 //kg->mgに変換
	raw := m * p / 100
	left := m - raw

	// 売れ残りをお惣菜にして売れた分を計算
	sidedish := left * q / 100
	left = left - sidedish

	// 最終的に売れ残った量を出力
	fmt.Println(float64(left) / 1000000) //表示のみmg->kgに変換
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/leftover
