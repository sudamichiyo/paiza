package main

import "fmt"

func main() {
	// 投球数を表す整数nを標準入力から入力
	var n int
	fmt.Scanf("%d", &n)

	// n回投球結果(strike または ball)を標準入力から入力
	pitching := make([]string, 0, n)
	for i := 0; i < n; i++ {
		var p string
		fmt.Scanf("%s", &p)
		pitching = append(pitching, p)
	}

	// 投球結果を出力(strikeの場合はstrike!と出力)
	scount := 0
	for _, pitch := range pitching {
		if pitch == "strike" {
			scount++
			fmt.Println(pitch + "!")
			fmt.Println(scount)
		} else {
			fmt.Println(pitch)
		}
	}
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/umpire_06
