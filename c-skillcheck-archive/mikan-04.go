package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 仕分ける重さの区切りを表す整数n, みかんの個数を表す整数mを標準入力から入力
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	// みかんを仕分ける箱の個数pを標準入力から入力
	var p int
	fmt.Scanf("%d", &p)

	// p個の箱に書かれた整数k1, k2, ... kpを空白区切りで標準入力から入力
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	k := strings.Split(scanner.Text(), " ")

	// string型のスライスをint型のスライスに変換
	intk := make([]int, len(k))
	for i, v := range k {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("変換エラー:", err)
		}
		intk[i] = num
	}

	// m行みかんの重さを標準入力から入力
	mikanwieght := make([]int, 0, m)
	for i := 0; i < m; i++ {
		scanner.Scan()
		w, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("変換エラー:", err)
		}
		mikanwieght = append(mikanwieght, w)
	}

	// 箱の整数がnの倍数のもののみにする
	var nk []int
	for _, v := range intk {
		if v%n == 0 {
			nk = append(nk, v)
		}
	}

	// みかんの仕分けを行う(絶対値の小さい方に入れる, 絶対値が等しい場合は大きい数字の箱に入れる)
	result := make([]int, 0, m)
	var inboxindex int
	for i, w := range mikanwieght {
		minabs := Abs(nk[0] - mikanwieght[i])
		for i := 0; i < len(nk); i++ {
			abs := Abs(nk[i] - w)
			if minabs >= abs {
				minabs = abs
				inboxindex = i
			}
		}
		result = append(result, nk[inboxindex])
	}

	// みかんの仕分け先の箱に書かれた数字を表す整数を出力
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

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/mikan_04
