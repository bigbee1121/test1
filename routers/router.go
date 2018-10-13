package routers

import (
	"dddd/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.InsertFilter("/article",beego.BeforeExec,BeforExecfunc)
    beego.Router("/", &controllers.MainController{} ,"get:ShowGet")

    beego.Router("/login", &controllers.UserController{},"get:ShowLogin;post:HandleLogin")

    beego.Router("/register",&controllers.UserController{},"get:ShowRegister;post:HandleRegister")
    //文章列表访问
	beego.Router("/article/showArticlelist",&controllers.ArticleController{},"get:ShowArticleList")

    beego.Router("/article/addArticle",&controllers.ArticleController{},"get:ShowAddArticle;post:HandleAddArticle")

    beego.Router("/article/showArticleDetail",&controllers.ArticleController{},"get:ShowArticleDetail")

    beego.Router("/article/UpdateArticle",&controllers.ArticleController{},"get:ShowUpdateArticle;post:HandleUpdate")

	beego.Router("/article/DleteAticle",&controllers.ArticleController{},"get:DeleteAticle")

	beego.Router("/article/addArticleType",&controllers.ArticleController{},"get:ShowArticleType;post:HandleArticleType")

	beego.Router("/logout",&controllers.UserController{},"get:Logout")

	beego.Router("/article/DeleteType",&controllers.ArticleController{},"get:DeleteType")
}
var BeforExecfunc = func (ctx *context.Context){
	username := ctx.Input.Session("username")
	if username == nil{
		ctx.Redirect(302,"/login")
		return
	}
}