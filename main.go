package main

import (
	"fmt"
)

const secret = "abc" // read-only

type Os int

const ( // iotaを指定すると自動で連番が付与される
	Mac     Os = iota + 1 // 1
	Windows               // 2
	Linux                 // 3
)

var ( // 一括で宣言できる
	ikkatsu_i int
	ikkatsu_s string
	ikkatsu_b bool
)

func main() {
	// var i int // funcの外だとvarを使うしかない
	// var i int = 2 // 初期化も可能
	// var i = 2 // 型推論も効く
	i := 1
	ui := uint16(2)
	fmt.Println(i)
	fmt.Printf("i: %v %T\n", i, i)                         // %v: value, %T: Type
	fmt.Printf("i: %[1]v %[1]T, ui: %[2]v %[2]T\n", i, ui) // 番号で指定できる

	f := 1.23456
	s := "hello"
	b := true
	fmt.Printf("f: %[1]v %[1]T\n", f)
	fmt.Printf("s: %[1]v %[1]T\n", s)
	fmt.Printf("b: %[1]v %[1]T\n", b)

	pi, title := 3.14, "Go" // 複数定義も可能
	fmt.Printf("pi: %v, title: %v\n", pi, title)

	x := 10
	y := 1.23
	z := float64(x) + y // 型変換は明示的に
	fmt.Printf("z: %v\n", z)

	// secret = "a" // Error("cannot assign to secret")
	fmt.Println(secret)

	fmt.Printf("Mac: %v, Windows: %v, Linux: %v\n", Mac, Windows, Linux)

	// 値の変更
	i = 2
	fmt.Printf("i: %v\n", i)
	i += 1
	fmt.Printf("i: %v\n", i)
	i *= 2
	fmt.Printf("i: %v\n", i)

	// 一括宣言
	fmt.Println(ikkatsu_i, ikkatsu_s, ikkatsu_b)
}
