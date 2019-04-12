package main

import (
	"fmt"
	"rens/infra/algo"
)

func test_after_shuffle() {
	fmt.Println(algo.AfterShuffle(int64(10), int64(100*100)))
}
