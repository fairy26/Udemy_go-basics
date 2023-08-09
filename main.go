package main

import (
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	flags := log.Lshortfile // ログに行数を含める
	warnLogger := log.New(io.MultiWriter(file, os.Stderr), "WARN: ", flags)
	errorLogger := log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", flags)

	warnLogger.Printf("warning A")
	errorLogger.Fatalln("critical error") // エラー内容を書き込んだ後にプログラムを強制的に終了する
}
