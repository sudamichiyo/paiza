package main

import "fmt"

func main() {
	// 合計の投球数nを標準入力から入力
	var n int
	fmt.Scanf("%d", &n)

	// n回投球の結果を表す文字列s("strike" or "ball")を標準入力から入力
	pitching := make([]string, 0, n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Scanf("%s", &s)
		pitching = append(pitching, s)
	}

	// ストライク、ボールの回数を管理する変数を定義
	scount := 0
	bcount := 0

	// 投球の結果によってコールを格納
	call := make([]string, 0, n)
	for _, p := range pitching {
		if p == "strike" { // 投球がストライクの場合
			scount++
			if scount > 2 { // ストライクが3回溜まった場合にアウト
				call = append(call, "out!")
			} else { // ストライクが3回未満の場合
				call = append(call, "strike!")
			}
		} else { // 投球がボールの場合
			bcount++
			if bcount > 3 { // ボールが4回溜まった場合にフォアボール
				call = append(call, "fourball!")
			} else { // ボールが4回未満の場合
				call = append(call, "ball!")
			}
		}
	}

	// 適切なコール( "strike!", "ball!", "out!", "fourball!" のいずれか)を出力
	for _, c := range call {
		fmt.Println(c)
	}

}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/umpire_01
