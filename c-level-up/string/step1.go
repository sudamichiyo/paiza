package main

import (
	"fmt"
	"math"
)

func main() {
	//  整数nを標準入力から入力
	var n int
	fmt.Scanf("%d", &n)

	// n個の整数を標準入力から入力
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	// 整数の文字列としての長さを計算し改行区切りで出力
	for i, _ := range a {
		var m float64
		if a[i] == 0 {
			m = 0.0
		} else {
			m = math.Log10(float64(a[i]))
		}
		len := math.Floor(m) + 1
		fmt.Println(len)
	}
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_level_up_problems/c_rank_string_step1
