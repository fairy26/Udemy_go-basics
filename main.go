package main

import (
	"fmt"
	"unsafe"
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

	var ui1 uint16
	fmt.Printf("memory address of ui1: %p\n", &ui1) // &でポインタに
	var ui2 uint16
	fmt.Printf("memory address of ui2: %p\n", &ui2)
	var p1 *uint16
	fmt.Printf("value of p1: %v\n", p1) // nilで初期化されている
	p1 = &ui1
	fmt.Printf("value of p1: %v\n", p1)
	fmt.Printf("size of p1: %d[bytes]\n", unsafe.Sizeof(p1)) // ポインタは8bytes
	fmt.Printf("memory address of p1: %p\n", &p1)            // "ダブルポインタ", "ポインタのポインタ"
	fmt.Printf("value of ui1(dereference): %v\n", *p1)
	*p1 = 1
	fmt.Printf("value of ui1: %v\n", ui1)

	var pp1 **uint16 = &p1
	fmt.Printf("value of pp1: %v\n", pp1)
	fmt.Printf("size of pp1: %d[bytes]\n", unsafe.Sizeof(pp1))
	fmt.Printf("memory address of pp1: %p\n", &pp1)
	fmt.Printf("value of p1(dereference): %v\n", *pp1)
	fmt.Printf("value of ui1(dereference): %v\n", **pp1)
	**pp1 = 10
	fmt.Printf("value of ui1: %v\n", ui1)

	ok, result := true, "A"
	fmt.Printf("memory address of result: %p\n", &result)
	if ok {
		// result := "B"   // 1: このスコープのみ (shadowing)
		result = "B"    // 2
		println(result) // 1: B, 2: B
		fmt.Printf("memory address of result: %p\n", &result)
	} else {
		result := "C"
		println(result)
	}
	println(result) // 1: A, 2: B
}
