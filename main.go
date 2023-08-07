package main

import (
	"errors"
	"fmt"
	"os"
)

var ErrCustom = errors.New("not found")

func main() {
	err01 := errors.New("something wrong")
	fmt.Printf("%[1]p, %[1]T, %[1]v\n", err01)
	// 0x14000096230, *errors.errorString, something wrong
	// errorの型はponter
	// ref. https://cs.opensource.google/go/go/+/refs/tags/go1.20.7:src/errors/errors.go;l=61

	// func Error() string {} を実装していると error interface を満たしているとみなされる
	fmt.Println(err01.Error())
	fmt.Println(err01) // error interfaceに準拠しているかもチェックしているため、そのまま手強くできる

	err02 := errors.New("something wrong")
	fmt.Println(err01 == err02) // アドレスの番地を比較している
	// false

	// errorのラップ; 既存のerrorに付加情報を与える
	err0 := fmt.Errorf("add info: %w", errors.New("original error")) // %wで既存のerrorを指定
	fmt.Printf("%[1]p, %[1]T, %[1]v\n", err0)
	// 0x14000060020, *fmt.wrapError, add info: original error
	fmt.Println(errors.Unwrap(err0))        // original error
	fmt.Printf("%T\n", errors.Unwrap(err0)) // *errors.errorString

	err1 := fmt.Errorf("add info: %v", errors.New("original error"))
	fmt.Println(err1)        // add info: original error
	fmt.Printf("%T\n", err1) // *errors.errorString
	// != *fmt.wrapError
	fmt.Println(errors.Unwrap(err1)) // <nil>
	// errors.errorString には Unwrap は実装されていないのでnilを返す

	// centinel error; Errから始まる規定のerror
	// ref. https://cs.opensource.google/go/go/+/master:src/os/error.go%3Bl=16
	err2 := fmt.Errorf("in repository layer: %w", ErrCustom)
	fmt.Println(err2)
	// in repository layer: not found

	err2 = fmt.Errorf("in service layer: %w", err2)
	fmt.Println(err2)
	// in service layer: in repository layer: not found

	// wrapされたerror(err2)がcentinel error(ErrCustom)と一致しているかを直接は調べられない
	// Unwrapする必要があるが、errors#Is でどこかの階層で一致しているかを調べることができる！
	if errors.Is(err2, ErrCustom) {
		fmt.Println("matched")
	}

	file := "dummy.txt"
	err3 := fileChecker(file)
	if err3 != nil {
		if errors.Is(err3, os.ErrNotExist) {
			fmt.Printf("%v file not found\n", file)
		} else {
			fmt.Println("unknown error")
		}
	}
	// dummy.txt file not found
}

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		// err: os.ErrNotExist
		return fmt.Errorf("in checker: %w", err)
	}
	defer f.Close()
	return nil
}
