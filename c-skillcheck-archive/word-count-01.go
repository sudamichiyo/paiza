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

	//　入力した文字列を半角スペース区切りでスライスに格納
	inputs := strings.Split(scanner.Text(), " ")

	// 文字列を出力
	for _, input := range inputs {
		fmt.Println(input)
	}
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/word-count_01
