package main

import (
	"fmt"
	"time"
)

func main() {
	i := -2

	if i == 0 {
		fmt.Println("zero")
	} else if i > 0 {
		fmt.Println("positive")
	} else {
		fmt.Println("negative")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for { // 無限ループ
	// 	fmt.Println("working")
	// 	time.Sleep(2 * time.Second)
	// }

	i = 0
	for {
		if i > 3 {
			break
		}
		fmt.Println(i)
		i += 1
		time.Sleep(300 * time.Millisecond)
	}

loop:
	for i := 0; i < 10; i++ {
		switch i {
		case 2:
			continue
		case 3:
			continue
		case 8:
			break loop // !
		default:
			fmt.Printf("%v ", i)
		}

	}
	fmt.Println("")

	items := []item{
		{price: 10.},
		{price: 20.},
		{price: 30.},
	}
	for _, i := range items {
		// i はコピーが生成されている
		i.price *= 1.1
	}
	fmt.Printf("%+v\n", items) // 値が変更されていない
	// [{price:10} {price:20} {price:30}]

	// 直接書き換える場合はindexで指定する
	for i := range items {
		items[i].price *= 1.1
	}
	fmt.Printf("%+v\n", items)
	// [{price:11} {price:22} {price:33}]
}

type item struct {
	price float32
}
