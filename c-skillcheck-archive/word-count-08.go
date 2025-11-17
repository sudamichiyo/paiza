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

	//単語をキー、単語の出現回数をバリューとするマップを作成
	count := make(map[string]int, len(inputs))

	// 単語の存在を確認し、新出単語の場合は単語リストに単語を追加する
	for _, word := range inputs {
		if !contains08(wordlist, word) { //新出単語の場合
			wordlist = append(wordlist, word)
		}
		// 単語の出現回数をカウント
		count[word]++
	}

	// 単語とその出現回数を表示
	for _, w := range wordlist {
		fmt.Println(w, count[w])
	}

}

func contains08(slice []string, word string) bool {
	for _, v := range slice {
		if v == word {
			return true
		}
	}
	return false
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/word-count_07
