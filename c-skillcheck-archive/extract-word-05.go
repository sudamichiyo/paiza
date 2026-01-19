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

	// 抽出データ格納用のスライス
	var result []string

	// 文字列データsに含まれる全ての開始タグと終了タグで囲まれた文字列を抽出する
	copy := s //処理用に文字列sをコピー
	stag := strings.Index(copy, taga)
	etag := strings.Index(copy, tagb)

	// 開始タグが見つからなくなるまで繰り返す
	for stag > -1 {
		cutdata := copy[stag+len(taga) : etag]
		if cutdata == "" {
			result = append(result, "<blank>")
		} else {
			result = append(result, cutdata)
		}

		// 終了タグを含めた前の文字列をカット
		copy = copy[etag+len(tagb):]
		stag = strings.Index(copy, taga)
		etag = strings.Index(copy, tagb)
	}

	// 結果を出力
	for _, cutstr := range result {
		fmt.Println(cutstr)
	}
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/extract_word_05
