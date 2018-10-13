package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"dddd/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) ShowRegister(){
	this.TplName = "register.html"
}

func (this *UserController)HandleRegister(){
	//抓取数据
	 Username := this.GetString("username")
	 Pwd := this.GetString("password")
	 //校验数据
	 if Username == ""|| Pwd == ""{
	 	this.Data["errmsg"] = "注册数据不完整！"
	 	//beego.Info("数据注册不完整！")
	 	this.TplName = "register.html"
	 	return
	 }
	 //处理数据
	 o := orm.NewOrm()
	 var user models.User

	 user.Name = Username
	 user.Password = Pwd

	 o.Insert(&user)

	 //this.Ctx.WriteString("注册成功！")
	 this.Redirect("/login",302)
	 //this.TplName="login.html"

}

func(this *UserController)ShowLogin(){
	userName := this.Ctx.GetCookie("username")
	if userName == ""{
		this.Data["userName"] = ""
		this.Data["checked"] = ""

	}else{
		this.Data["userName"] = userName
		this.Data["checked"] = "checked"
	}
	this.TplName = "login.html"
}

func(this *UserController)HandleLogin(){
	 Username := this.GetString("username")
	 Pwd := this.GetString("password")

	 if Username == ""||Pwd == ""{
	 	this.Data["errmsg"] = "注册数据不完整！"
	 	this.TplName = "login.html"
	 }

	 o:= orm.NewOrm()

	 var user models.User

	 user.Name = Username
	 err := o.Read(&user,"Name")
	 if err != nil{
	 	this.Data["errmsg"] = "用户不存在！"
	 	this.TplName = "login.html"
	 	return
	 }
	 if user.Password != Pwd{
	 	this.Data["errmsg"] = "密码错误！"
	 	this.TplName = "login.html"
	 	return
	 }
	 data := this.GetString("remember")
	 beego.Info(data)

	 if data == "on"{
		 this.Ctx.SetCookie("username",Username,100)
	 }else{
		 this.Ctx.SetCookie("username",Username,-1)
	 }

	this.SetSession("username",Username)
	 //this.Ctx.SetCookie("Username",Username,100)
this.Redirect("/article/showArticlelist",302)
	 //this.Ctx.WriteString("登录成功！")
}
func(this*UserController)Logout(){
	//删除session
	this.DelSession("username")
	//跳转登录页面
	this.Redirect("/login",302)
}