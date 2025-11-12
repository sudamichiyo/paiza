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

	// スライス内の単語が新出であればその単語を出力、既出ならalready_beenを出力
	// 単語リストを格納するスライスを定義
	wordlist := []string{}

	// 単語の存在を確認し、新出単語の場合は単語リストに単語を追加する
	for _, word := range inputs {
		if contains03(wordlist, word) { //既出単語の場合
			fmt.Println("already_been")
		} else { // 新出単語の場合
			wordlist = append(wordlist, word)
			fmt.Println(word)
		}
	}

}

func contains03(slice []string, word string) bool {
	for _, v := range slice {
		if v == word {
			return true
		}
	}
	return false
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/word-count_03
