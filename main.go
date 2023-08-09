package main

// 関数名を右クリック→「Go: Generate Unit Tests For Function」
// テーブルドリブンのテストが自動生成される
func Add(x, y int) int {
	return x + y
}

func Devide(x, y int) float32 {
	if y == 0 {
		return 0.
	}
	return float32(x) / float32(y)
}

func main() {
	// x, y := 3, 5
	// fmt.Printf("%v, %v\n", Add(x, y), Devide(x, y))
	// カバレッジを計算するためにコメントアウトで一旦対象外に
}

// テストを実行する
// ```bash
// $ go test -v .
// === RUN   TestAdd
// === RUN   TestAdd/1+2=3
// --- PASS: TestAdd (0.00s)
//     --- PASS: TestAdd/1+2=3 (0.00s)
// PASS
// ok      go-basics       0.272s
// ```

// カバレッジを計算する
// ```bash
// $ go test -v -cover -coverprofile=coverage.out .
// === RUN   TestAdd
// === RUN   TestAdd/1+2=3
// === RUN   TestAdd/2+3=5
// --- PASS: TestAdd (0.00s)
//     --- PASS: TestAdd/1+2=3 (0.00s)
//     --- PASS: TestAdd/2+3=5 (0.00s)
// === RUN   TestDevide
// === RUN   TestDevide/1/2=0.5
// --- PASS: TestDevide (0.00s)
//     --- PASS: TestDevide/1/2=0.5 (0.00s)
// PASS
//         go-basics       coverage: 75.0% of statements
// ok      go-basics       0.122s  coverage: 75.0% of statements
// ```

// coverageを可視化する
// ```bash
// $ go tool cover -html=coverage.out
// ```
