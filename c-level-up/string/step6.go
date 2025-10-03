package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 時刻を表す長さ5の文字列sを標準入力から入力
	var s string
	fmt.Scanf("%s", &s)

	// ":"で区切って時と分に分けて変数に格納
	t := strings.Split(s, ":")

	// 30分後の時間を計算
	var h, m int
	var err error
	if h, err = strconv.Atoi(t[0]); err == nil {
		if m, err = strconv.Atoi(t[1]); err == nil {
			m += 30
			if m >= 60 {
				h += 1
				m %= 60
			}
		}
	}

	// "XX:XX"の形式で時刻を出力
	time := fmt.Sprintf("%02d", h) + ":" + fmt.Sprintf("%02d", m)
	fmt.Println(time)
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_level_up_problems/c_rank_string_step6
