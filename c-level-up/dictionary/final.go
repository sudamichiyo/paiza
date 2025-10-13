package main

import (
	"fmt"
	"sort"
)

func main() {
	// Aグループの人数p、Bグループの人数q、Cグループの人数rを標準入力から入力
	var p, q, r int
	fmt.Scanf("%d %d %d", &p, &q, &r)

	// p行Aグループの⚪︎番目の人がBグループの△番目の人に仕事を任せるかを標準入力から入力
	job := make([][]int, p)
	for i := range job {
		job[i] = make([]int, 2)
	}

	for i := 0; i < p; i++ {
		var a, b int
		fmt.Scanf("%d %d", &a, &b)
		job[i][0] = a
		job[i][1] = b
	}

	// q行Bグループのj番目の人がCグループのk番目の人に仕事を任せるかを標準入力から入力（Aグループから頼まれていないB->Cの余計な仕事も混ざっている可能性がある）
	for i := 0; i < q; i++ {
		var b, c int
		fmt.Scanf("%d %d", &b, &c)
		// B の値を持つすべての A 要素に対して C を追加する
		for j := range job {
			if job[j][1] == b {
				job[j] = append(job[j], c)
			}
		}
	}

	// 仕事をAグループの番号が小さい順にソート
	sort.Slice(job, func(i, j int) bool {
		x, y := job[i], job[j]
		// 0番目の値で昇順に比較
		return x[0] < y[0]
	})

	// 最終的にAグループのi番目の仕事をCグループの誰が担当することになるのかを半角スペース区切りで出力

	for i := range job {
		fmt.Println(job[i][0], job[i][2])
	}
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_level_up_problems/c_rank_dictionary_boss
