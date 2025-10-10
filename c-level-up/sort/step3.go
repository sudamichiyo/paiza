package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 正整数nを標準入力から入力
	var n int
	fmt.Scanf("%d", &n)

	// n個の数のペアを標準入力から入力して二次元スライスに格納
	pair := make([][]int, n)
	for i := 0; i < n; i++ {
		pair[i] = make([]int, 2)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < 2; j++ {
			var num int
			fmt.Scanf("%d", &num)
			pair[i][j] = num
		}
	}

	// りんごの数が大きい順にソート、ただしりんごの数が同じ場合はバナナの数が大きい順にソート
	for i := 0; i < len(pair)-1; i++ {
		for j := 0; j < len(pair)-(i+1); j++ {
			// りんごの数を比較
			if pair[j][0] < pair[j+1][0] {
				tmp := pair[j]
				pair[j] = pair[j+1]
				pair[j+1] = tmp
			} else if pair[j][0] == pair[j+1][0] {
				// バナナの数を比較
				if pair[j][1] < pair[j+1][1] {
					tmp := pair[j]
					pair[j] = pair[j+1]
					pair[j+1] = tmp
				}
			}
		}
	}

	// int型のスライスをstring型に変換
	spair := make([][]string, n)
	for i := range spair {
		spair[i] = make([]string, 2)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < 2; j++ {
			s := strconv.Itoa(pair[i][j])
			spair[i][j] = s
		}
	}

	// 結果のペアを改行区切りで出力
	for _, p := range spair {
		fmt.Println(strings.Join(p, " "))
	}
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_level_up_problems/c_rank_sort_step3
