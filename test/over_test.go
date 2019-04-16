package test

import (
	"fmt"
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	"testing"
)

// 基准测试代码,使用Benchmark 开头，第一个字母B大写
func BenchmarkUpdateForLock(b *testing.B) {
	g := GoodsSigned{}
	g.EnvelopeNo=ksuid.New().Next().String()
	g.RemainQuantity = 100000
	g.RemainAmount = decimal.NewFromFloat(100)
	_,err:= db.Insert(g)
	if err != nil {
	fmt.Println(err)
		return
	}

	for i:= 0;i< b.N ;i++  {
		UpdateForLock(g.Goods)
	}
}

// 基准测试 ：无符号字段类型+直接更新  TODO: 扣减失败
func BenchmarkUpdateForUnsigned(b *testing.B) {
	g :=GoodsSigned{}
	g.EnvelopeNo = ksuid.New().Next().String()
	g.RemainQuantity = 100000
	g.RemainAmount = decimal.NewFromFloat(1000000)
	_,err:=db.Insert(g)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i <b.N ; i++ {
		UpdateForUnsigned(g.Goods)
	}
}