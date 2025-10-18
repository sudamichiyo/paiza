package main

import "fmt"

func main() {
	// 13x >= 10000よりx >= 10000/13を求める
	// そのうち最小の自然数13xを求める
	x := 0
	if 10000%13 != 0 { //13で割り切れなかった場合
		x = 10000/13 + 1
	} else { //13で割り切れる場合
		x = 10000 / 13
	}
	answer := 13 * x

	fmt.Println(answer)
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_level_up_problems/c_rank_simulation_step1
