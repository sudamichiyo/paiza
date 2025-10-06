package main

import (
	"fmt"
	"time"
)

var layout = "15:04"

func main() {
	// 工事が続く週の数nを標準入力から入力
	var n int
	fmt.Scanf("%d", &n)

	// n行工事が始まる時刻t、工事が続く時間h、分mを標準入力から入力
	start := make([]string, 0, n)
	hours := make([]int, 0, n)
	minutes := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var s string
		var h, m int
		fmt.Scanf("%s %d %d", &s, &h, &m)
		start = append(start, s)
		hours = append(hours, h)
		minutes = append(minutes, m)
	}

	// 時刻を文字列からtime.Time型に変換
	t := make([]time.Time, 0, n)
	for i := range start {
		t = append(t, stringToTime(start[i]))
	}

	// 工事の終了時刻を計算
	for i := range t {
		t[i] = t[i].Add(time.Duration(hours[i])*time.Hour + time.Duration(minutes[i])*time.Minute)
	}

	// n行工事が終わる時刻を出力
	time := make([]string, 0, n)
	for i := range t {
		time = append(time, timeToString(t[i]))
		fmt.Println(time[i])
	}
}

func timeToString(t time.Time) string {
	str := t.Format(layout)
	return str
}

func stringToTime(str string) time.Time {
	t, _ := time.Parse(layout, str)
	return t
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_level_up_problems/c_rank_string_boss
