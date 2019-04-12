package main

import (
	"fmt"
	"rens/infra/algo"
)

func main1() {
	// test_simple_rand()
}

func test_simple_rand() {
	count, amout := int64(10), int64(100*100)
	for i := int64(0); i < count; i++ {
		// fmt.Println("最大发出的值是：", amout, count)
		num := count - i
		x := algo.SimpleRand(num, amout)
		// fmt.Println("生产的随机书包：", x)
		amout -= x
		fmt.Print(float64(x)/float64(100), ",")
	}
	fmt.Println()
}
