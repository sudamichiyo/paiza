package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 整数nを標準入力から入力
	var n int
	fmt.Scanf("%d", &n)

	// 数列aを標準入力から入力
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputs := strings.Split(scanner.Text(), " ")

	// 文字列型のスライスinputsをint型のスライスに変換
	sequence := make([]int, 0, len(inputs))
	for _, s := range inputs {
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Error converting %s: %v\n", s, err)
			continue
		}
		sequence = append(sequence, i)
	}

	// 数列aが等差数列かチェックして等差数列である場合は"Yes"をそうでない場合は"No"を出力
	d := sequence[1] - sequence[0] //公差
	an := sequence[0] + (len(sequence)-1)*d
	if an == sequence[len(sequence)-1] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

	// 数列aが等比数列かチェック等比数列である場合は"Yes"をそうでない場合は"No"を出力
	r := float64(sequence[1]) / float64(sequence[0])
	bn := float64(sequence[0]) * math.Pow(float64(r), float64(len(sequence)-1))
	if bn == float64(sequence[len(sequence)-1]) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

// 問題URL:https://paiza.jp/works/mondai/c_rank_skillcheck_archive/sequence_check
