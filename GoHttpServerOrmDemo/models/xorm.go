package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
	"xorm.io/core"
)

var Engine *xorm.Engine

//xorm框架初始化
func init()  {

	var err error

	//创建orm引擎
	Engine, err = xorm.NewEngine("mysql","root:123456@tcp(127.0.0.1:3306)/demo")

	if err != nil {
		fmt.Println(err)
		return
	}

	//连接测试
	if err := Engine.Ping(); err!=nil{
		fmt.Println(err)
		return
	}

	//打印sql
	Engine.ShowSQL(true)

	//配置时区
	if location, err := time.LoadLocation("Asia/Shanghai"); err == nil {
		Engine.TZLocation = location
	}

	//设置连接池空闲数大小
	Engine.SetMaxIdleConns(5)
	//设置最大打开连接数
	Engine.SetMaxOpenConns(5)

	//名称映射规则主要是负责结构体到表名和表中字段名称的映射
	//SameMapper 支持结构体名称和对应的表名称以及结构体field名称与对应的表字段名称相同的命名； *
	Engine.SetColumnMapper(core.SameMapper{})
	Engine.SetTableMapper(core.GonicMapper{})

	err = Engine.CreateTables(&HotelDict{}) //SameMapper规则创建表

	if err != nil {
		fmt.Println("CreateTables errors:", err)
	}
}




