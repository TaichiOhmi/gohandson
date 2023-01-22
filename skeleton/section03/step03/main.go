package main

import (
	"fmt"
	"math/rand"
	"time"
)

// TODO: string型をベースにしてsuit型を定義する
// ユーザ定義型
type suit string

const (
	suitHeart   suit = "♥"
	suitClub    suit = "♣"
	suitDiamond suit = "◆"
	// TODO: ♠を表す定数suitSpadeを定義する
	// 名前付き定数
	suitSpade suit = "♠"
)

type card struct {
	suit suit
	// TODO: int型のnumberフィールドを定義する
	number int
}

func main() {

	suits := []suit{
		suitHeart,
		suitClub,
		suitDiamond,
		suitSpade,
	}
	// fmt.Println(suits)
	// [♥ ♣ ◆ ♠]

	// 山札を作る
	all := make([]card, 0, 13*4) // 長さが0で、容量は13*4の52枚
	for _, s := range suits {    // ハート、クラブ、ダイヤ、スペードの４種類を一番外でループ。
		for n := 2; n <= 14; n++ { // 2~14までの数字をループ
			all = append(all, card{
				// 各suitsにおいて、2~14の数値がallに追加される。
				// TODO: マークをセットする
				// 構造体リテラル
				suit:   s,
				number: n,
			})
		}
	}

	// 乱数の種をセットする
	t := time.Now().UnixNano()
	rand.Seed(t)

	// 山札をシャッフルさせる
	rand.Shuffle(len(all), func(i, j int) {
		all[i], all[j] = all[j], all[i]
	})

	// 手札
	cards := all[:5]
	all = all[5:]

	// 手札を表示させる
	// cardsの要素を1つずつ取り出し変数cに入れる
	for _, c := range cards {
		// "♥ "のように出力する
		fmt.Print(c.suit, " ")
		switch c.number {
		case 11:
			fmt.Println("J")
		case 12:
			fmt.Println("Q")
		case 13:
			fmt.Println("K")
		case 14:
			fmt.Println("A")
		default:
			// TODO: 番号を改行ありで出力する
			// フィールドを参照している。
			fmt.Println(c.number)
		}
	}
}
