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

	// スライス内の単語が新出であればその単語を単語リストに格納
	// 単語リストを格納するスライスを定義
	wordlist := []string{}

	//単語をキー、1をバリューとするマップを作成（続きの問題で1の部分には単語の出現回数が入る）
	count := make(map[string]int, len(inputs))

	// 単語の存在を確認し、新出単語の場合は単語リストに単語を追加する
	for _, word := range inputs {
		if !contains05(wordlist, word) { //新出単語の場合
			wordlist = append(wordlist, word)
			count[word] = 1
		}
	}

	// 単語リストを表示
	for _, w := range wordlist {
		fmt.Println(w)
	}

	// 単語の種類の数だけ1を出力
	for _, v := range count {
		fmt.Println(v)
	}

}

func contains05(slice []string, word string) bool {
	for _, v := range slice {
		if v == word {
			return true
		}
	}
	return false
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/word-count_05
