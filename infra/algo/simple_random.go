package algo

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	min int64 = 1 // 最小1分钱
)

/*
	TODO : 该算法中存在超卖的情况
	可能遇到的问题：
	1. 每次生成的随机数都一样，
		原因：需要设置种子数，如时间戳


	该版本的问题：
		生成的金额数是从大到小，越往后越小
*/

// 简单随机算法
// 红包的数量，红包金额  单位：分  类型: int64
func SimpleRand(count, amount int64) int64 {
	// 最大可调度金额 = 红包总金额-最小的分配金额
	var max, x int64
	if count == 10 {
		max = amount - min*count
	} else {
		max = amount
	}
	rand.Seed(time.Now().UnixNano()) // 设置种子数
	fmt.Println("最大发：", max)
	if max > 0 {
		x = rand.Int63n(max) + min
		fmt.Println("当前随机数：", x) // 存在问题：如果当前余额很小，产生的随机数可能等于其本身，导致之后的数为负数，发生超卖的情况
	} else {
		return min
	}
	return x
}
