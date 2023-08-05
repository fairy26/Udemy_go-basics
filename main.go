package main

import (
	"fmt"
	"unsafe"
)

type Task struct {
	Title    string
	Estimate int
}

func main() {
	task1 := Task{
		Title:    "Learning Golang",
		Estimate: 3,
	}
	task1.Title = "Learning Go"
	fmt.Printf("%[1]T, %+[1]v, %v\n", task1, task1.Title) // %vに+をつけると構造体のフィールド名も出力
	// main.Task, {Title:Learning Go Estimate:3}, Learning Go

	var task2 Task = task1 // 実体は別のメモリ領域
	task2.Title = "new"
	fmt.Printf("task1: %v, task2: %v\n", task1.Title, task2.Title)
	// task1: Learning Go, task2: new

	task1p := &Task{
		Title:    "Learning concurrency",
		Estimate: 2,
	}
	fmt.Printf("task1p: %T, %+v, %v\n", task1p, *task1p, unsafe.Sizeof(task1p))
	// task1p: *main.Task, {Title:Learning concurrency Estimate:2}, 8

	// (*task1p).Title = "Changed"
	task1p.Title = "Changed" // dereferenceの(*__)は省略できる！
	fmt.Printf("task1p: %+v\n", *task1p)
	// task1p: {Title:Changed Estimate:2}

	var task2p *Task = task1p
	task2p.Title = "Changed by Task2"
	fmt.Printf("task1: %+v\n", *task1p)
	fmt.Printf("task2: %+v\n", *task2p)
	// task1: {Title:Changed by Task2 Estimate:2}
	// task2: {Title:Changed by Task2 Estimate:2} // pointerだから共有している

	task1.extendEstimate()
	fmt.Printf("task1 value receiver: %+v\n", task1.Estimate)
	// task1 value receiver: 3 // 元のtask1には影響を与えない

	// (&task1).extendEstimatePointer() // pointerを取得する(&__)は省略できる！
	(&task1).extendEstimatePointer()
	fmt.Printf("task1 value receiver: %+v\n", task1.Estimate)
	// task1 value receiver: 13

}

// reciever
func (task Task) extendEstimate() {
	// 受けとった構造体のコピーに対して操作を行う
	task.Estimate += 10
}

func (task *Task) extendEstimatePointer() {
	task.Estimate += 10 // dereference(*__)の省略形
}
