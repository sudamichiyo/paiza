package main

import "fmt"

func main() {
	// 整数nを標準入力から入力
	var n int
	fmt.Scanf("%d", &n)

	// 整数a, bを標準入力から入力
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	// 手続きの回数をカウントする変数を定義
	count := 0

	// 初期値をセット
	paiza := 1
	kyoko := 1

	// 以下の処理を京子の数がnより大きくなるまで繰り返す
	for kyoko <= n {
		// パイザ君は自分の数のa倍を京子の数に足す
		kyoko += paiza * a
		count++
		// 京子は自分の数をbで割った余りをパイザ君の数に足す
		paiza += kyoko % b
	}

	// 手続きが終わった時のパイザ君の操作回数を出力
	fmt.Println(count)
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_level_up_problems/c_rank_simulation_step2
