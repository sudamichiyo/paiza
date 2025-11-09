package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 半角スペースで区切られた文字列を標準入力から入力
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputs := strings.Split(scanner.Text(), " ")

	// 出現した単語を格納する配列
	wordlist := []string{}
	// 単語の出現回数を管理するマップ
	count := make(map[string]int, len(inputs))

	// 新出単語であれば単語リストの配列に格納
	for _, word := range inputs {
		if !wordContains(wordlist, word) {
			wordlist = append(wordlist, word)
		}

		// 単語の出現回数をカウント
		count[word]++
	}

	// 単語リストの格納順に単語と出現回数を半角スペース区切りで出力
	for _, word := range wordlist {
		fmt.Println(word, count[word])
	}
}

func wordContains(slice []string, key string) bool {
	for _, v := range slice {
		if v == key {
			return true
		}
	}
	return false
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/word-count_00
