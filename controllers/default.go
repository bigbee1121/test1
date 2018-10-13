package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"dddd/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["data"] = "vim"
	c.TplName = "test.html"
}

func (p *MainController)Post(){
	p.TplName = "test.html"

	p.Data["data"] = "haihaihaih"

}
func (c *MainController)ShowGet(){
	o := orm.NewOrm()
	//
	//var user models.User
	//
	//user.Name = "heima"
	//user.PassWord = "chuanzhi"
	//id,err := o.Insert(&user)
	//if err!=nil{
	//	beego.Info(err)
	//	return
	//}
	//beego.Info(id)

	//查询操作
	//var user models.User
	//user.Id = 1
	//
	//err := o.Read(&user)
	//if err != nil{
	//	beego.Error("查询失败")
	//}
	////返回结果
	//beego.Info(user)

	//更新
//var user models.User
//user.Id = 1
//err := o.Read(&user)
//if err != nil{
//	beego.Error("更新数据不存在！")
//}
//user.Name  ="shanghai"
//count,err := o.Update(&user)
//if err != nil{
//	beego.Error("更新失败")
//}
//beego.Info(count)


	var user models.User
	user.Id = 1
	user.Id = 2
	count,err := o.Delete(&user)
	if err != nil{
		beego.Error("删除失败！")
	}
	beego.Info(count)
	c.Data["data"] = "上海"
	c.TplName = "test.html"
}
