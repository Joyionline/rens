package main

import (
	"fmt"
	"rens/infra/algo"
)

func Test_after_shuffle() {
	fmt.Println(algo.AfterShuffle(int64(10), int64(100*100)))
}
