package main

import "fmt"

func main() {
	//参加人数n, 正整数m, 正整数kを標準入力から入力
	var n, m, k int
	fmt.Scanf("%d %d %d", &n, &m, &k)

	// 各参加者のm個の数を標準入力から入力
	numbers := make([][]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scanf("%d", &numbers[i][j])
		}
	}

	// 各参加者のポイントを計算
	result := make([]int, 0, n)
	for i := range numbers {
		result = append(result, countNumber(numbers[i], k))
	}

	// 各参加者のポイントを出力
	for _, count := range result {
		fmt.Println(count)
	}

}

// 参加者の手元の数字から指定された数kがいくつ含まれているかを数える関数
func countNumber(slice []int, k int) int {
	count := 0
	for _, v := range slice {
		if v == k {
			count++
		}
	}
	return count
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_level_up_problems/c_rank_for_boss
