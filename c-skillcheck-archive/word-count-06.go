package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 半角スペースで区切られた長さNの文字列を標準入力から入力
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	// 半角スペースで区切ってスライスに格納
	inputs := strings.Split(scanner.Text(), " ")

	// 文字列sを標準入力から入力
	scanner.Scan()
	s := scanner.Text()

	// 文字列sが1行目で半角スペース区切りで与えられた文字列の何番目に登場するかを出力
	// (条件に当てはまる文字列が複数あった場合、最初にその文字列が現れたのが何番目であるかを出力)
	for i, word := range inputs {
		if word == s {
			fmt.Println(i)
			return
		}
	}
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/word-count_06
