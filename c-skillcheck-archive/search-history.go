package main

import (
	"fmt"
)

func main() {
	// 検索ワードの数を表す整数nを標準入力から入力
	var n int
	fmt.Scanf("%d", &n)

	// 検索ワードwをスライスに格納
	// ただし、検索ワードwが以前に入力されたことがある場合は履歴中のwを削除して履歴の先頭にwを追加
	// 検索ワードwが以前に入力されたことがない場合は履歴の先頭にwを追加
	history := []string{}
	for i := 0; i < n; i++ {
		word := ""
		fmt.Scanf("%s", &word)
		ok, index := contains(history, word)
		if ok { //wが存在している場合
			// wを削除
			history = append(history[:index], history[index+1:]...)
			// wを末尾に追加
			history = append(history, word)
		} else { //wが存在していない場合
			// wを末尾に追加
			history = append(history, word)
		}
	}

	// 検索ワードをn個入力した後の検索履歴を出力(スライスを逆順に出力)
	for i := len(history) - 1; i >= 0; i-- {
		fmt.Println(history[i])
	}
}

func contains(s []string, w string) (ok bool, index int) {
	for i, v := range s {
		if w == v {
			return true, i
		}
	}
	return false, -1
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/search_history
