package algo

import (
	"math/rand"
	"time"
)

func DoubleRandom(count, amount int64) int64 {
	if count == 1 {
		return amount
	}
	// 计算最大可调度金额
	max := amount - min*count
	rand.Seed(time.Now().UnixNano())
	// 一次随机，计算出一个种子金额作为基数
	seed := rand.Int63n(count*2) + 1
	n := max/seed + min
	x := rand.Int63n(n)
	return x
}
