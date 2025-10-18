package main

import "fmt"

func main() {
	//　パイザ君の体力Hを標準入力から入力
	var H int
	fmt.Scanf("%d", &H)

	// パイザ君とモンスターの攻撃回数, ダメージを数える変数を定義
	pattack := 0
	mattack := 0
	var pdamege []int
	var mdamege []int

	// パイザ君の体力Hが0以下になるまで攻撃を繰り返す
	for H > 0 {
		pattack++
		if pattack == 1 || pattack == 2 { // パイザ君の魔法は 1 回目と 2 回目に使うときにはダメージ 1
			mdamege = append(mdamege, -1)
		} else { // 3 回目以降の n 回目には、(モンスターから受けた (n - 1) 回目の攻撃のダメージ) + (モンスターから受けた (n - 2) 回目の攻撃のダメージ) のダメージを与える
			damege := pdamege[pattack-2] + pdamege[pattack-3]
			mdamege = append(mdamege, damege)
		}

		mattack++
		if mattack == 1 || mattack == 2 { // モンスターの魔法は 1 回目と 2 回目に使うときにはダメージ 1
			pdamege = append(pdamege, -1)
			H += -1
		} else { // モンスターの魔法は3 回目以降の n 回目には、 (パイザ君から受けた (n - 1) 回目の攻撃のダメージ) * 2 + (パイザ君から受けた (n - 2) 回目の攻撃のダメージ) のダメージを与える
			damege := mdamege[mattack-2]*2 + mdamege[mattack-3]
			pdamege = append(pdamege, damege)
			H += damege
		}
	}

	// モンスターの何回目の攻撃でパイザ君の体力が 0 以下になるかを出力
	fmt.Println(mattack)
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_level_up_problems/c_rank_simulation_boss
