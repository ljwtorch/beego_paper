package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	//设置数据库基本信息
	orm.RegisterDataBase("default", "mysql", dsn)
	//注册模型
	orm.RegisterModel(new(User), new(Article), new(ArticleType))
	// 生成表,第一个参数是别名，第2个参数是否强制更新（除非表结构发生改变），一般设置为false，若为true，则表数据会被清空！
	// 第3个参数是是否打印日志（打印sql语句）
	orm.RunSyncdb("default", false, true)
}

func TableName(str string) string {
	return beego.AppConfig.String("dbprefix") + str
}
