package model

import (
	"fmt"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

var database sqlbuilder.Database

func init()  {
	InitUpperDatabase()
}

func UpperDatabase() sqlbuilder.Database  {
	if database == nil {
		InitUpperDatabase()
	}
	return database
}


func InitUpperDatabase()  {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("严重错误：",err)
		}
	}()
	settings:=mysql.ConnectionURL{
		Database:"po",
		Host:"localhost",
		User:"root",
		Password:"joyiever",
	}
	db,err:=mysql.Open(settings)
	if err != nil {
		panic(err)
	}
	fmt.Println("数据库链接成功")
	database = db
}