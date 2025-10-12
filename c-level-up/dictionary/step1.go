package main

import "fmt"

func main() {
	// 人数nを標準入力から入力
	var n int
	fmt.Scanf("%d", &n)

	// n人の名前と財産を標準入力から入力
	m := make(map[string]int, n)
	for i := 0; i < n; i++ {
		var name string
		var assets int
		fmt.Scanf("%s %d", &name, &assets)
		m[name] = assets
	}

	// 人名sを標準入力から入力
	var person string
	fmt.Scanf("%s", &person)

	// 人名sに一致する財産を出力
	v, ok := m[person]
	if !ok {
		fmt.Println("名前がありません")
	} else {
		fmt.Println(v)
	}
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_level_up_problems/c_rank_dictionary_step1
