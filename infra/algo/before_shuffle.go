package algo

import "math/rand"

// 先洗牌算法
func BeforeShuffle(count, amout int64) int64 {
	if count == 1 {
		return amout
	}
	// 计算最大可调度金额
	max := amout - min*count

	// 生成红包种子金额序列
	seeds := make([]int64, 0)
	// 红包种子金额序列长度=3～1/2*count
	size := count / 2
	if size < 3 {
		size = 3
	}
	for i := int64(0); i < size; i++ {
		x := max / (i + 1)
		seeds = append(seeds, x)
	}
	// 从红包种子金额金额序列中随机选择一个作为随机基数
	idx := rand.Int63n(int64(len(seeds)))
	x := rand.Int63n(seeds[idx])
	return x + min
}
