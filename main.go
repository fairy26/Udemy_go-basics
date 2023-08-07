package main

import (
	"fmt"
	"unsafe"
)

type controller interface { // メソッドの一覧を定義する
	speedUp() int
	speedDown() int
}

type vehicle struct {
	speed        int
	machinePower int
}

type bycycle struct {
	speed      int
	humanPower int
}

func (v *vehicle) speedUp() int {
	v.speed += 10 * v.machinePower
	return v.speed
}

// ---
func (v *vehicle) speedDown() int {
	v.speed -= 5 * v.machinePower
	return v.speed
}

// ---
// interfaceに書かれたすべてのメソッドを実装しているtypeは
// 自動的に「そのinterfaceを実装している」とみなされる
// → そのinterfaceを引数にとる関数などが使える！

func (v *bycycle) speedUp() int {
	v.speed += 5 * v.humanPower
	return v.speed
}

func (v *bycycle) speedDown() int {
	v.speed -= 3 * v.humanPower
	return v.speed
}

func speedUpAndDown(c controller) {
	fmt.Printf("current speed: %v\n", c.speedUp())
	fmt.Printf("current speed: %v\n", c.speedDown())
}

func (v vehicle) String() string {
	// Sprintg: フォーマットしたstringを"返す"
	return fmt.Sprintf("vehicle current speed is %v (enginePower is %v)", v.speed, v.machinePower)
}

func main() {
	v := &vehicle{0, 5}
	// メソッドを呼び出す時にpointer receiverを使っているため、&でポインタを取得する
	speedUpAndDown(v) // 渡せる！

	b := &bycycle{0, 5}
	speedUpAndDown(b)

	// Stringer interface
	// ref. https://pkg.go.dev/golang.org/x/tools/cmd/stringer
	// fmtで出力するとき、内部で Stringer interface を満たすかチェックし、
	// その #String() を呼び出している
	// ; Stringer関数を実装すればfmtで出力できる
	fmt.Println(v)

	// any型; すべての型をとれる
	var i1 any
	var i2 interface{} // any型と全く一緒; 別表記 (0個のメソッドを実装する型)
	fmt.Printf("%[1]%v, %[1]T, %v\n", i1, unsafe.Sizeof(i1))
	fmt.Printf("%[1]%v, %[1]T, %v\n", i2, unsafe.Sizeof(i2))
	// %v, <nil>, 16
	checkType(i1) // nil
	i1 = 1
	checkType(i1) // int
	i1 = "hello"
	checkType(i1) // string

}

func checkType(i any) {
	switch i.(type) {
	case nil:
		fmt.Println("nil")

	case int:
		fmt.Println("int")

	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}
