package main

import "fmt"

func main() {
	// 整数nを標準入力から入力
	var n int
	fmt.Scanf("%d", &n)

	// 1からnまでの整数を1から順に表示
	// ただし、表示しようとしている数字が15の倍数のときにはFizzBuzz、3の倍数の時にはFizz、5の倍数の時にはBuzzをかわりに出力
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			fmt.Println("Fizz Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_sample/fizz-buzz
