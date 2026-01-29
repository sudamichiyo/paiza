package main

import (
	"fmt"
	"strings"
)

func main() {
	// 3つの正整数a, b, cを標準入力から文字列として入力
	var a, b, c string
	fmt.Scanf("%s", &a)
	fmt.Scanf("%s", &b)
	fmt.Scanf("%s", &c)

	// a, b, cをスライスに格納
	var number []string
	number = append(number, a)
	number = append(number, b)
	number = append(number, c)

	// a, b, cをハイフン(-)で区切って出力
	phonenumber := strings.Join(number, "-")
	fmt.Println(phonenumber)

}

// 問題URL:https://paiza.jp/works/mondai/rank_test_problems_c_0001/rank_test_problems_c_0001__1
