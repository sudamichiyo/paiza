package main

import "fmt"

func main() {
	// 仕分ける重さの区切りを表す整数n, みかんの個数を表す整数mを標準入力から入力
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	// みかんを仕分ける箱に書かれた整数k1, k2を標準入力から入力
	var k1, k2 int
	fmt.Scanf("%d %d", &k1, &k2)

	// m行みかんの重さを標準入力から入力
	mikanwieght := make([]int, 0, m)
	for i := 0; i < m; i++ {
		var w int
		fmt.Scanf("%d", &w)
		mikanwieght = append(mikanwieght, w)
	}

	// みかんの仕分けを行う(絶対値の小さい方に入れる, 絶対値が等しい場合は大きい数字の箱に入れる)
	result := make([]int, 0, m)
	for _, w := range mikanwieght {
		abs1 := Abs(k1 - w)
		abs2 := Abs(k2 - w)
		if abs1 >= abs2 {
			result = append(result, k2)
		} else {
			result = append(result, k1)
		}
	}

	// みかんの仕分け先の箱に書かれた数字を表す整数（k1 or k2）を出力
	for _, box := range result {
		fmt.Println(box)
	}
}

// 絶対値を求める関数
func Abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/mikan_03
