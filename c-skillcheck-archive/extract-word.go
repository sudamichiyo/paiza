package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 開始タグと終了タグを半角スペース区切りで標準入力から入力
	var taga, tagb string
	fmt.Scanf("%s %s", &taga, &tagb)

	// テキストデータを標準入力から入力
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	// データを処理するためにテキストデータをコピー
	copy := text
	// 処理結果格納用のスライス
	var result []string

	// テキストデータ中の開始タグと終了タグの出現位置を調べる
	stag := strings.Index(copy, taga)
	etag := strings.Index(copy, tagb)

	// 開始タグが見つからなくなるまで処理を繰り返す
	for stag > -1 {
		// 文字列を切り取り
		cutdata := copy[stag+len(taga) : etag]

		// 切り取った文字列をスライスに格納する
		if cutdata == "" {
			result = append(result, "<blank>")
		} else {
			result = append(result, cutdata)
		}

		// 処理用のテキストデータcopyの終了タグを含めその前をカットする
		copy = copy[etag+len(tagb):]

		// copyの開始タグと終了タグの出現位置を調べる
		stag = strings.Index(copy, taga)
		etag = strings.Index(copy, tagb)
	}

	//結果を出力
	for _, cutstr := range result {
		fmt.Println(cutstr)
	}
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/extract_word_00
