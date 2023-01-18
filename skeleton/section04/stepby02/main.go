package main

import "fmt"

type T int

func (t *T) M() {
	fmt.Printf("M: %d\n", int(*t))
}

func (t T) G() {
	fmt.Printf("G: %d\n", t)
}

func main() {
	ts := []T{10, 20, 30}

	ms := make([]func(), 0, len(ts))
	gs := make([]func(), 0, len(ts))

	for _, t := range ts {
		ms = append(ms, t.M)
		//ms = append(ms, (&t).M)

		//gs = append(gs, t.G)
		gs = append(gs, (&t).G)
	}

	for i, _ := range ts {
		ms[i]()
		// 30
		// 30
		// 30
	}
	for i, _ := range ts {
		gs[i]()
		// 10
		// 20
		// 30
	}
}
