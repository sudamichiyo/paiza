package main

import (
	"fmt"
	"strings"
)

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
	bcount := 0
	for _, pitch := range pitching {
		if pitch == "strike" { //strikeの場合
			scount++
			if scount > 2 {
				fmt.Println(strings.Replace(pitch, "strike", "out", 1) + "!")
			} else {
				fmt.Println(pitch + "!")
			}
		} else { //ballの場合
			bcount++
			if bcount > 3 {
				fmt.Println(strings.Replace(pitch, "ball", "fourball", 1) + "!")
			} else {
				fmt.Println(pitch + "!")
			}
		}
	}
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/umpire
