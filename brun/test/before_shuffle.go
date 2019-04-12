package main

import (
	"fmt"
	"rens/infra/algo"
)

func main() {
	test_before_shuffle()
	test_double_random()
}

func test_before_shuffle() {
	count, amount := int64(10), int64(100)
	var sum float64
	for i := int64(0); i < count; i++ {
		x := algo.BeforeShuffle(count, amount*100)
		sum += float64(x) / 100
		fmt.Print(float64(x)/100, ",")
	}
	fmt.Println("发出的红包总金额是：", sum)
}
