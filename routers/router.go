package routers

import (
	"beego_paper/controllers"
	"github.com/astaxie/beego"
	// 注意要导入下面的context包
	"github.com/astaxie/beego/context"
)

func init() {
	//过滤器函数，判断用户是否登录
	beego.InsertFilter("/article/*", beego.BeforeRouter, filterFun)

	// 路由设置
	// 注意：当实现了自定义的get请求方法，那么post及其其他请求将不会访问默认方法
	beego.Router("/register", &controllers.RegisterController{}, "get:TurnToRegister;post:HandleRegister")

	// 首先需要登录
	beego.Router("/login", &controllers.LoginAndLogoutController{}, "get:TurnToLogin;post:HandleLogin")
	beego.Router("/", &controllers.LoginAndLogoutController{}, "get:TurnToLogin;post:HandleLogin")
	beego.Router("/logout", &controllers.LoginAndLogoutController{}, "get:Logout")

	// 论文详情的路由无需添加参数
	beego.Router("/article/showArticleList", &controllers.ArticleController{}, "get:TurnToArticleList")
	beego.Router("/article/showArticleDetail", &controllers.ArticleController{}, "get:TurnToArticleDetail")
	beego.Router("/article/addArticle", &controllers.ArticleController{}, "get:TurnToAddArticle;post:HandleAddArticle")
	beego.Router("/article/updateArticle", &controllers.ArticleController{}, "get:TurnToUpdateArticle;post:HandleUpdateArticle")
	beego.Router("/article/deleteArticle", &controllers.ArticleController{}, "get:HandleDeleteArticle")

	// 论文类型操作
	beego.Router("/article/addArticleType", &controllers.ArticleTypeController{}, "get:TurnToAddArticleType;post:HandleAddArticleType")
	beego.Router("/article/deleteArticleType", &controllers.ArticleTypeController{}, "get:HandleDeleteArticleType")
}

//过滤器限制登录
var filterFun = func(ctx *context.Context) {
	userName := ctx.Input.Session("userName")
	if userName == nil {
		ctx.Redirect(302, "/")
	}
}
