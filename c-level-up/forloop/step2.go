package main

import "fmt"

func main() {
	// 正整数nを標準入力から入力
	var n int
	fmt.Scanf("%d", &n)

	// n個の整数を標準入力から入力
	numbers := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scanf("%d", &num)
		numbers = append(numbers, num)
	}

	// 与えられた整数の中に7が一つでもあればflagをtrueにする
	flag := false
	for _, v := range numbers {
		if v == 7 {
			flag = true
			break
		}
	}

	// 整数の中に7が一つでも含まれていればYES, 一つも含まれていなければNOを出力
	if flag {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}

// 問題URL:https://paiza.jp/works/mondai/c_rank_level_up_problems/c_rank_for_step2
