package test

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/core/errors"
	"github.com/shopspring/decimal"
	"github.com/tietang/dbx"
)

var db *dbx.Database

func init() {
	settings := dbx.Settings{
		DriverName: "mysql",
		Host:       "localhost",
		User:       "root",
		Password:   "joyiever",
		Database:   "po0",
		Options: map[string]string{
			"parseTime": "true",
		},
	}
	var err error
	db, err = dbx.Open(settings)
	if err != nil {
		fmt.Println("Error", err)
	}
	db.RegisterTable(&GoodsSigned{},"goods")
	db.RegisterTable(&GoodsUnsigned{},"goods_unsigned")
}


// 事务锁方案
func UpdateForLock(g Goods)  {
	// 通过db.tx 函数构建事务锁代码块
	err:=db.Tx(func(runner *dbx.TxRunner) error {
		// 1.锁定需要修改的资源，即修改的数据行
		query := "select * from goods where envelope_no = ? for update" // 实际开发中用需要查询的字段代替*，提高性能
		out := &GoodsSigned{}
		ok, err := runner.Get(out, query, g.EnvelopeNo)
		if !ok || err != nil {
			return err
		}
		// 2.计算剩余金额和数量
		subAmount := decimal.NewFromFloat(0.01)
		remainAmount := out.RemainAmount.Sub(subAmount)
		remainQuantity := out.RemainQuantity - 1
		// 3. 执行更新
		updatesql := "update goods " +
			" set remain_amount = ?, remain_quantity = ?" +
			" where envelope_no = ? "
		_, row, err := db.Execute(updatesql, remainAmount, remainQuantity,g.EnvelopeNo)
		if err != nil {
			return err
		}
		if row < 1 {
			return errors.New("库存扣减失败")
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

// 数据库无符号类型+直接更新方案
func UpdateForUnsigned(g Goods)  {
	updatesql := "update goods_unsigned " +
			" set remain_amount=remain_amount - ?, " +
			" remain_quantity=remain_quantity - ? " +
			" where envelope_no = ?"
	_,row,err:=db.Execute(updatesql,0.01,1,g.EnvelopeNo)
	if err!= nil {
		fmt.Println(err)
	}
	if row < 1 {
		fmt.Println("扣减失败")
	}
}