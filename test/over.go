package test

import (
	"fmt"
	"resk-pilot/test"

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
	db.RegisterTable(&test.GoodsSigned{},"goods")
	db.RegisterTable(&test.GoodsUnsigned{},"goods_unsigned")
}


// 事务锁方案
func UpdateForLock(g Goods)  {
	// 通过db.tx 函数构建事务锁代码块
	db.Tx(func(run *dbx.TxRunner) error {
		// 锁定需要修改的资源，即修改的数据行
		query := "select * from goods where envelope_no=? for update"
		out := &test.GoodsSigned{}
		runner.Get()
	})
}