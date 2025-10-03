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
	time := strings.Split(s, ":")

	// 時、分の順番で改行区切りで出力(入力値の十のくらいが0である場合には一の位のみ出力)
	if h, err := strconv.Atoi(time[0]); err == nil {
		fmt.Println(h)
	}
	if m, err := strconv.Atoi(time[1]); err == nil {
		fmt.Println(m)
	}
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_level_up_problems/c_rank_string_step5
