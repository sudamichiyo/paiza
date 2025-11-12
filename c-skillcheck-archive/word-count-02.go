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

	// 文字列を半角スペース区切りでスライスに格納
	inputs := strings.Split(scanner.Text(), " ")

	// スライスの中にredが存在するかを確認、あればYes、なければNoを出力する
	if contains02(inputs, "red") {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func contains02(slice []string, word string) bool {
	for _, v := range slice {
		if v == word {
			return true
		}
	}
	return false
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/word-count_02
