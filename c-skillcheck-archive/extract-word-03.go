package main

import (
	"fmt"
	"strings"
)

func main() {
	// 半角スペースで区切られた2つの<>で囲まれた開始タグと終了タグを標準入力から入力
	var taga, tagb string
	fmt.Scanf("%s %s", &taga, &tagb)

	// 抽出処理を行う文字列データsを標準入力から入力
	var s string
	fmt.Scanf("%s", &s)

	// 文字列データsに含まれる開始タグと終了タグの開始位置を空白区切りで出力
	stag := strings.Index(s, taga)
	etag := strings.Index(s, tagb)
	stag = stag + 1
	etag = etag + 1
	fmt.Println(stag, etag)
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/extract_word_03
