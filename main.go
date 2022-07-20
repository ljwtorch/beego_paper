package main

import (
	_ "beego_paper/controllers/utils"
	"beego_paper/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	// 下划线'_'表示只执行routers的init()函数
	_ "beego_paper/routers"
)

func init() {
	models.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true
}

func HandlePrePage(pageIndex int) int {
	pageIndex -= 1
	return pageIndex
}

func HandleNextPage(pageIndex int) int {
	pageIndex += 1
	return pageIndex
}

func main() {
	//映射函数的执行必须在run之前被调用
	_ = beego.AddFuncMap("ShowPrePage", HandlePrePage)
	_ = beego.AddFuncMap("ShowNextPage", HandleNextPage)

	//在开发环境下，您可以使用以下指令来开启查询调试模式
	orm.Debug = true
	beego.Run()
}
