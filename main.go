package main

import (
	"fmt"
)

func main() {
	// +---
	// | array
	// +---
	var a1 [3]int
	var a2 = [3]int{10, 20, 30}
	a3 := [...]int{10, 20} // ...で要素数を自動で数えてくれる
	fmt.Println(a1, a2, a3)
	fmt.Printf("%v, %v\n", len(a3), cap(a3)) // 2, 2
	fmt.Printf("%T, %T\n", a2, a3)           // [3]int, [2]int
	// → 型が違うので代入したりはできない

	// +---
	// | slice
	// +---
	var s1 []int
	s2 := []int{}
	fmt.Printf("s1: %[1]T, %[1]v %v %v\n", s1, len(s1), cap(s1)) // []int, [] 0 0
	fmt.Printf("s2: %[1]T, %[1]v %v %v\n", s2, len(s2), cap(s2)) // []int, [] 0 0
	fmt.Println(s1 == nil)                                       // true
	fmt.Println(s2 == nil)                                       // false

	// 要素の追加
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("s1: %[1]T, %[1]v %v %v\n", s1, len(s1), cap(s1)) // []int, [1 2 3] 3 3
	s3 := []int{4, 5, 6}
	s1 = append(s1, s3...)                                       // ...で展開してappendする
	fmt.Printf("s1: %[1]T, %[1]v %v %v\n", s1, len(s1), cap(s1)) // []int, [1 2 3 4 5 6] 6 6

	s4 := make([]int, 0, 2)
	fmt.Printf("s4: %[1]T, %[1]v %v %v\n", s4, len(s4), cap(s4)) // []int, [] 0 2
	s4 = append(s4, 1, 2, 3, 4)                                  // appendではあらかじめ定めたcapを超えることができる
	fmt.Printf("s4: %[1]T, %[1]v %v %v\n", s4, len(s4), cap(s4)) // []int, [1 2 3 4] 4 4

	s5 := make([]int, 4, 6)                            // len分0で初期化
	fmt.Printf("s5: %v %v %v\n", s5, len(s5), cap(s5)) // [0 0 0 0] 4 6
	s6 := s5[1:3]
	s6[1] = 10                                         // index [1, 3); {s5[1], s5[2]}
	fmt.Printf("s5: %v %v %v\n", s5, len(s5), cap(s5)) // [0 0 10 0] 4 6 // こっちも変わる
	fmt.Printf("s6: %v %v %v\n", s6, len(s6), cap(s6)) // [0 10] 2 5
	s6 = append(s6, 2)
	fmt.Printf("s5: %v %v %v\n", s5, len(s5), cap(s5))          // [0 0 10 2] 4 6 // 元の値が2に置き換わっている
	fmt.Printf("s6 appended: %v %v %v\n", s6, len(s6), cap(s6)) // [0 10 2] 3 5
	// → sliceを切り取って使用するとメモリを共有してしまう

	sc6 := make([]int, len(s5[1:3]))
	fmt.Printf("s5 source of copy: %v %v %v\n", s5, len(s5), cap(s5))      // [0 0 10 2] 4 6
	fmt.Printf("sc6 dst copy before: %v %v %v\n", sc6, len(sc6), cap(sc6)) // [0 0] 2 2
	copy(sc6, s5[1:3])                                                     // メモリを共有したくない時はcopyを使う
	fmt.Printf("sc6 dst copy after: %v %v %v\n", sc6, len(sc6), cap(sc6))  // [0 10] 2 2
	sc6[1] = 12
	fmt.Printf("s5: %v %v %v\n", s5, len(s5), cap(s5))     // [0 0 10 2] 4 6 // 変わらず！
	fmt.Printf("sc6: %v %v %v\n", sc6, len(sc6), cap(sc6)) // [0 12] 2 2

	// メモリの共有を部分的に
	s5 = make([]int, 4, 6)
	fs6 := s5[1:3:3]                                       // メモリを共有する最大のindexを最後の要素で指定(index+1)
	fmt.Printf("s5: %v %v %v\n", s5, len(s5), cap(s5))     // [0 0 0 0] 4 6
	fmt.Printf("fs6: %v %v %v\n", fs6, len(fs6), cap(fs6)) // [0 0] 2 2
	fs6[0] = 6
	fs6[1] = 7
	fs6 = append(fs6, 8)
	fmt.Printf("s5: %v %v %v\n", s5, len(s5), cap(s5))     // [0 6 7 0] 4 6 // 最大のindex=2までは共有されているがappend分はされず
	fmt.Printf("fs6: %v %v %v\n", fs6, len(fs6), cap(fs6)) // [6 7 8] 3 4
	s5[3] = 9
	fmt.Printf("s5: %v %v %v\n", s5, len(s5), cap(s5))     // [0 6 7 9] 4 6
	fmt.Printf("fs6: %v %v %v\n", fs6, len(fs6), cap(fs6)) // [6 7 8] 3 4 // 逆も影響されない

	// +---
	// | map
	// +---
	var m1 map[string]int
	m2 := map[string]int{}
	fmt.Printf("%v %v\n", m1, m1 == nil) // map[] true
	fmt.Printf("%v %v\n", m2, m2 == nil) // map[] false
	m2["A"] = 10
	m2["B"] = 20
	m2["C"] = 0
	fmt.Printf("%v %v %v\n", m2, len(m2), m2["A"]) // map[A:10 B:20 C:0] 3 10
	delete(m2, "A")
	fmt.Printf("%v %v %v\n", m2, len(m2), m2["A"]) // map[B:20 C:0] 2 0
	// 存在しない0と値の0の判別（返り値の2番目）
	v, ok := m2["A"]
	fmt.Printf("%v %v\n", v, ok) // 0 false
	v, ok = m2["C"]
	fmt.Printf("%v %v\n", v, ok) // 0 true

	for key, value := range m2 {
		fmt.Printf("%v %v\n", key, value)
	}
	// B 20
	// C 0
	// ※内部でハッシュマップを使っているので、順番は必ず同じとは限らない！
}
