package db

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{
	Type:     "mysql",
	User:     "root",
	Password: "123456",
	Host:     "127.0.0.1:3306",
	Name:     "xingyeblog",
}

var db *xorm.Engine // 全局的engine

func init() {
	var err error

	db, err = xorm.NewEngine(DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DatabaseSetting.User,
		DatabaseSetting.Password,
		DatabaseSetting.Host,
		DatabaseSetting.Name))

	if err != nil {
		fmt.Println("mysql connect err:", err)
	}
	fmt.Println("数据库连接成功")
	db.ShowSQL(true)

	err = db.Ping()
	if err != nil {
		fmt.Println("db connect error:", err.Error())
	}

	//10m keep connection 每十分钟测试一下数据库是否断开
	timer := time.NewTicker(time.Minute * 10)
	go func(db *xorm.Engine) {
		for _ = range timer.C {
			err = db.Ping()
			if err != nil {
				fmt.Println("db connect error:", err.Error())
			}
		}
	}(db)
}
