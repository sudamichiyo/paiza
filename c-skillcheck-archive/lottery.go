package main

import "fmt"

func main() {
	// 当選番号b, 購入した宝くじの数nを標準入力から入力
	var b int
	var n int
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &n)

	// 宝くじ券の番号を標準入力から入力
	lotnums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &lotnums[i])
	}

	// 当選結果を管理するマップを定義
	results := make(map[int]string, n)

	// 当選結果を確認
	for _, num := range lotnums {
		if num == b { // 1等(当選番号と一致する番号)の場合
			results[num] = "first"
		} else if num+1 == b || num-1 == b { // 前後賞(当選番号の ±1 の番号)の場合
			results[num] = "adjacent"
		} else if num%10000 == b%10000 { // 2等(当選番号と下 4 桁が一致する番号)の場合
			results[num] = "second"
		} else if num%1000 == b%1000 { // 3等(当選番号と下 3 桁が一致する番号)の場合
			results[num] = "third"
		} else { //その他(はずれ)の場合
			results[num] = "blank"
		}
	}

	// 結果を出力
	for _, num := range lotnums {
		fmt.Println(results[num])
	}
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/lottery
