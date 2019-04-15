package algo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSimpleRand(t *testing.T) {
	count, amout := int64(0), int64(10000)
	remain := amout
	sum := int64(0)
	for i := int64(0); i < count; i++ {
		x := SimpleRand(count-i, remain)
		remain -= x
		sum += x
	}
	Convey("简单随机算法", t, func() {
		So(sum, ShouldEqual, amout)
	})
}
