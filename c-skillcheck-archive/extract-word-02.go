package main

import "fmt"

func main() {
	// 半角スペースで区切られた2つの<>で囲まれた開始タグと終了タグを標準入力から入力
	var taga, tagb string
	fmt.Scanf("%s %s", &taga, &tagb)

	// 抽出処理を行う文字列データsを標準入力から入力
	var s string
	fmt.Scanf("%s", &s)

	// 開始タグと終了タグの開始位置を標準入力から入力
	var aindex, bindex int
	fmt.Scanf("%d %d", &aindex, &bindex)
	aindex = aindex - 1
	bindex = bindex - 1

	// 文字列sに含まれる開始タグと終了タグで囲まれた部分を出力
	// もし、囲まれた文字列がない場合は<blank>という文字を出力
	cutdata := s[aindex+len(taga) : bindex]
	if cutdata == "" {
		fmt.Println("<blank>")
	} else {
		fmt.Println(cutdata)
	}
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/extract_word_02
