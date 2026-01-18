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

	// 文字列データsに含まれる全ての開始タグと終了タグの開始位置を空白区切りで出力
	copy := s //処理用に文字列sをコピー
	stag := strings.Index(copy, taga)
	etag := strings.Index(copy, tagb)
	totalstag := stag
	totaletag := etag

	// 開始タグが見つからなくなるまで繰り返す
	for stag > -1 {
		fmt.Println(totalstag+1, totaletag+1)
		preendindex := totaletag

		// 終了タグを含めた前の文字列をカット
		copy = copy[etag+len(tagb):]
		stag = strings.Index(copy, taga)
		etag = strings.Index(copy, tagb)

		// インデックスを足す
		totalstag = preendindex + len(tagb) + stag
		totaletag = preendindex + len(tagb) + etag
	}
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/extract_word_01
