package main

import (
	"fmt"
	"sort"
)

func main() {
	// 正整数a, b, cを標準入力から文字列として入力
	var a, b, c string
	fmt.Scanf("%s %s %s", &a, &b, &c)

	// 文字を配列に格納
	slice := make([]string, 0, 3)
	slice = append(slice, a)
	slice = append(slice, b)
	slice = append(slice, c)

	// a+bとb+aを比べる(他のペアでも同様に比べる)
	sort.Slice(slice, func(i, j int) bool {
		return slice[i]+slice[j] > slice[j]+slice[i]
	})

	// 3つをつなげてできる整数の内、最大のものを出力
	result := ""
	for _, s := range slice {
		result += s
	}

	fmt.Println(result)
}

// 問題URL:https://paiza.jp/works/mondai/rank_test_problems_c_0003/rank_test_problems_c_0003__3
