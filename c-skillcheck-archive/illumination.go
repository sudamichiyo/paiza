package main

import "fmt"

func main() {
	// 整数n(LEDの個数), m(信号の数)を標準入力から入力
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	// LEDをスライスで定義
	led := make([]int, n)

	// LEDライトの明るさの合計
	sum := 0

	// 信号(LEDの番号と明るさ)をm回標準入力から入力
	num := make([]int, 0, m)
	brightness := make([]int, 0, m)
	for i := 0; i < m; i++ {
		var n, b int
		fmt.Scanf("%d %d", &n, &b)
		num = append(num, n-1)
		brightness = append(brightness, b)
	}

	for i := 0; i < m; i++ {
		// ledに値を上書き
		led[num[i]] = brightness[i]

		// LEDライトが同時に点灯しているかチェック
		for _, b := range led {
			check := 1
			check *= b
			if check == 0 {
				sum = 0
				break
			} else {
				sum += b
			}
		}

		// n個のLEDが初めて同時に点灯したときにループを抜ける
		if sum > 0 {
			break
		}
	}

	// n個のLEDが初めて同時に点灯したときの明るさの合計を出力
	fmt.Println(sum)
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/illumination
