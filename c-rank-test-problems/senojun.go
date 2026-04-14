package main

import (
	"fmt"
	"sort"
)

func main() {
	// 生徒数nを標準入力から入力
	var n int
	fmt.Scanf("%d", &n)

	// 生徒の身長aと生徒の名前sを標準入力から入力
	type Person struct {
		Name   string
		Height int
	}
	people := []Person{}

	var s string
	var h int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %s", &h, &s)
		people = append(people, Person{s, h})
	}

	// 身長が高い順にソート
	sort.Slice(people, func(i, j int) bool {
		return people[i].Height > people[j].Height
	})

	// 身長が高い生徒から順に改行区切りで名前を出力
	for _, person := range people {
		fmt.Println(person.Name)
	}
}

// 問題URL:https://paiza.jp/works/mondai/rank_test_problems_c_0001/rank_test_problems_c_0001__4
