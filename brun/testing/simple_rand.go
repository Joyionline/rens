package main

import (
	"fmt"
	"rens/infra/algo"
)

func main() {
	count, amout := int64(10), int64(100)
	for i := int64(0); i < count; i++ {
		x := algo.SimpleRand(count, amout*100)
		fmt.Print(float64(x)/float64(100), ",")
	}
	fmt.Println()
}
