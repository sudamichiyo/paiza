package main

import "fmt"

func main() {
	// 投球数を表す整数nを標準入力から入力
	var n int
	fmt.Scanf("%d", &n)

	// strike or ball という文字列を改行区切りでn回出力
	for i := 0; i < n; i++ {
		fmt.Println("strike or ball")
	}
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/umpire_03
