package main

import (
	"fmt"
	"sync"
)

// +-----
// | 4. goroutineの注意点
// +-----
func main() {
	var wg sync.WaitGroup
	s := []int{1, 2, 3}

	// for _, i := range s {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		fmt.Println(i)
	// 	}()
	// }
	// wg.Wait()
	// 3
	// 3
	// 3
	// 起動に時間がかかるので最後のiが参照される

	for _, i := range s {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
	// 3
	// 1
	// 2
	// goroutineに与えることができるが、順番は担保されない
}

// +-----
// | 3. トレース: 逐次処理 vs 並列処理
// +-----
// func main() {
// 	// tracerのlogファイルの作成
// 	f, err := os.Create("trace.out")
// 	if err != nil {
// 		log.Fatalln("Error: ", err) // プログラムの強制終了
// 	}

// 	defer func() {
// 		if err := f.Close(); err != nil {
// 			log.Fatalln("Error: ", err)
// 		}
// 	}()

// 	// tracerの開始/終了
// 	if err := trace.Start(f); err != nil {
// 		log.Fatalln("Error: ", err)
// 	}
// 	defer trace.Stop()

// 	// contextの作成 (ref. 後の講座で)
// 	ctx, t := trace.NewTask(context.Background(), "main")
// 	defer t.End()

// 	// ロジカルコア数の表示
// 	fmt.Println("The number of logical CPU Cores:", runtime.NumCPU())

// 	// 逐次処理
// 	// task(ctx, "Task1")
// 	// task(ctx, "Task2")
// 	// task(ctx, "Task3")

// 	// 並列処理
// 	var wg sync.WaitGroup
// 	wg.Add(3)
// 	go cTask(ctx, &wg, "Task1")
// 	go cTask(ctx, &wg, "Task2")
// 	go cTask(ctx, &wg, "Task3")
// 	wg.Wait()
// 	fmt.Println("main func finish")
// }

// func task(ctx context.Context, name string) {
// 	defer trace.StartRegion(ctx, name).End() // method chainでは最後のmethod(#End())のみdeferされる
// 	time.Sleep(time.Second)
// 	fmt.Println(name)
// }
// func cTask(ctx context.Context, wg *sync.WaitGroup, name string) {
// 	defer trace.StartRegion(ctx, name).End() // method chainでは最後のmethod(#End())のみdeferされる
// 	defer wg.Done()
// 	time.Sleep(time.Second)
// 	fmt.Println(name)
// }

// ```bash
// $ go tool trace trace.out
// ```

// +-----
// | 2. goroutineのjoin
// +-----

// func main() {
// 	var wg sync.WaitGroup

// 	wg.Add(1) // count += 1
// 	go func() {
// 		defer wg.Done() // count -= 1
// 		fmt.Println("goroutine invoked")
// 	}()

// 	wg.Wait() // til count == 0
// 	fmt.Println("num of working goroutines:", runtime.NumGoroutine())
// 	// 1
// 	fmt.Println("main func finished")
// }

// +-----
// | 1. groutineの起動
// +-----

// func main() {
// 	go func() {
// 		fmt.Println("goroutine invoked")
// 	}()

// 	fmt.Println("num of working goroutines:", runtime.NumGoroutine())
//  // 2
// 	fmt.Println("main func finished")
// 	// goroutineが起動するまでに時間がかかるため、先にmainが終了する
// 	// ("goroutine invoked"が表示されない)
// }
