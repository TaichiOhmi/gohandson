package main

import "fmt"

// typeで型を定義
type T int

// 定義した型のレシーバを指定し、メソッドを定義
func (t *T) M() {
	fmt.Println("method")
}

func main() {
	var t T
	p := &t
	p.M()
	t.M()
	// ここでは、t.M()と(&t).M()は同じように解釈される。
}
