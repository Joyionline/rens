package algo

import (
	"math/rand"
	"time"
)

// 二倍均值算法
func DoubleAverage(count, amout int64) int64 {
	if count == 1 {
		return amout
	}
	// 计算最大可调用金额
	max := amout - min*count
	// 计算最大可用平均值
	avg := max - count
	// 防止出现0的情况，在二倍的基础上加上最小值
	avg2 := 2*avg + min
	//随机红包金额序列，把二倍均值作为随机的最大数
	rand.Seed(time.Now().UnixNano())
	x := rand.Int63n(avg2) + min
	return x
}
