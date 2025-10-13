package main

import (
	"fmt"
	"sort"
)

func main() {
	// 人数nを標準入力から入力
	var n int
	fmt.Scanf("%d", &n)

	// n人の名前を標準入力から入力
	p := make(map[string]int, n)
	for i := 0; i < n; i++ {
		var name string
		fmt.Scanf("%s", &name)
		// 全ての人のダメージを0で初期化
		p[name] = 0
	}

	// 攻撃回数mを標準入力から入力
	var m int
	fmt.Scanf("%d", &m)

	// m行攻撃された人名とダメージを標準入力から入力
	for i := 0; i < m; i++ {
		var name string
		var damage int
		fmt.Scanf("%s %d", &name, &damage)
		_, ok := p[name]
		if !ok {
			fmt.Println("名前が存在しません")
		} else {
			p[name] += damage
		}
	}

	// 人名を辞書順に整列
	// キーをスライスに格納する
	keys := make([]string, 0, n)
	for k := range p {
		keys = append(keys, k)
	}

	// 辞書順にソート
	sort.Strings(keys)

	// ダメージを名前の辞書順に出力（一度も攻撃を受けていない場合のダメージは0）（出力はダメージのみ）
	for _, k := range keys {
		fmt.Println(p[k])
	}
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_level_up_problems/c_rank_dictionary_step3
