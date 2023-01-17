package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 山札を作ります
	all := make([]int, 0, 13)
	for n := 2; n <= 14; n++ {
		all = append(all, n)
	}

	// fmt.Println(all)
	// [2 3 4 5 6 7 8 9 10 11 12 13 14]

	// 乱数の種を現在時刻で取得
	t := time.Now().UnixNano()
	// 乱数の種を設定
	rand.Seed(t)

	// 山札をシャッフルさせる
	rand.Shuffle(len(all), func(i, j int) {
		all[i], all[j] = all[j], all[i]
	})

	// TODO: 山札の前方5枚を手札としcardsに入れる
	cards := all[:5] // 0~4枚目までの5枚
	// TODO: 6枚目以降を新しい山札とする
	all = all[5:] // 5枚目以降

	// 手札を表示させます
	for i, n := range cards {
		fmt.Printf("%d番目: ", i+1)
		// この行以降で、n が、case の場合、
		switch n {
		case 11:
			fmt.Println("J")
		case 12:
			fmt.Println("Q")
		case 13:
			fmt.Println("K")
		case 14:
			fmt.Println("A")
		default:
			fmt.Println(n)
		}
	}
}
