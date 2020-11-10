package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var (
		d time.Duration
		f float64
	)
	// Rust &(参照渡し)　と同じだね。
	// 関数の引数の定義が、、、*なんだよね。。w
	// →ポインタを示す意味だから辻褄わあうけどw
	flag.DurationVar(&d, "dur", 1*time.Second, "duration flag")
	flag.Float64Var(&f, "float", 0.1, "float flag")
	flag.Parse()
	fmt.Println(d, f)
}
