package main

import "fmt"

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

	// 人名sを標準入力から入力
	var person string
	fmt.Scanf("%s", &person)

	// sが受けたダメージを出力（一度も攻撃を受けていない場合のダメージは0）
	fmt.Println(p[person])
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_level_up_problems/c_rank_dictionary_step2
