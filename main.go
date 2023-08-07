package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func funcDefer() {
	// 複数のdifferは下から実行される
	defer fmt.Println("main func final-finish")
	defer fmt.Println("main func semi-finish")
	fmt.Println("hello world")
}

func trimExtension(files ...string) []string { // ...string で可変長のsliceを扱える
	out := make([]string, 0, len(files))
	for _, f := range files {
		out = append(out, strings.TrimSuffix(f, ".csv"))
	}
	return out
}

func fileChecker(name string) (string, error) {
	f, err := os.Open(name)
	if err != nil {
		return "", errors.New("file not found")
	}
	defer f.Close()
	return name, nil
}

func addExt(f func(file string) string, name string) {
	fmt.Println(f(name))
}

func multiply() func(int) int {
	return func(i int) int {
		return i * 1000
	}
}

func countUp() func(int) int {
	count := 0 // global変数とは違い、値をこのスコープに閉じ込めることができる
	return func(n int) int {
		count += n
		return count
	}
}

func main() {
	funcDefer()

	files := []string{"file1.csv", "file2.csv", "file3.csv"}
	fmt.Println(trimExtension(files...))

	name, err := fileChecker("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name) // $ touch file.txt

	// 無名関数
	i := 1
	func(i int) {
		fmt.Println(i)
	}(i) // 引数に与えると即座に実行してくれる

	f1 := func(i int) int {
		return i + 1
	}
	fmt.Println(f1(i))

	// 無名関数を引数にもつ関数
	f2 := func(name string) string {
		return name + ".csv"
	}
	addExt(f2, "file1")

	// 無名関数を返り値にもつ関数
	f3 := multiply()
	fmt.Println(f3(2))

	// closure
	f4 := countUp()
	for i := 1; i <= 5; i++ {
		v := f4(2)
		fmt.Printf("%v\n", v)
	}
}
