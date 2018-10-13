package controllers

import (
	"github.com/astaxie/beego"
	"path"
	"time"
	"dddd/models"
	"github.com/astaxie/beego/orm"
	"math"

)

type ArticleController struct {
	beego.Controller
}


func (this *ArticleController)ShowArticleList(){
	username := this.GetSession("username")
	if username == nil{
		this.Redirect("/login",302)
		return
	}


	typename := this.GetString("select")
	o:= orm.NewOrm()
	qs := o.QueryTable("Article")

	var count int64

	if typename == ""{
		count,_ = qs.RelatedSel("ArticleType").Count()
	}else{
		count,_=qs.RelatedSel("ArticleType").Filter("ArticleType__Tname",typename).Count()
	}

	beego.Info(typename)





	pageSize := 2
	pageCount :=math.Ceil(float64(count) / float64(pageSize))
	//获取页码
	pageIndex,err:= this.GetInt("pageIndex")
	if err != nil{
		pageIndex = 1
	}

	//获取数据
	//作用就是获取数据库部分数据,第一个参数，获取几条,第二个参数，从那条数据开始获取,返回值还是querySeter
	//起始位置计算
	start := (pageIndex - 1)*pageSize


	var articles []models.Article
	//_,err := qs.All(&articles)
	//if err != nil{
	//	beego.Info("查询数据错误")
	//}
	//this.Data["articles"] = articles
	//this.TplName = "index.html"
	if typename ==""{
		qs.Limit(pageSize,start).RelatedSel("ArticleType").All(&articles)
	}else{
		qs.Limit(pageSize,start).RelatedSel("ArticleType").Filter("ArticleType__Tname",typename).All(&articles)
	}

	var articletypes []models.ArticleType
	o.QueryTable("ArticleType").All(&articletypes)


	Username := this.GetSession("username")
	this.Data["username"] = Username.(string)
	this.Data["articletypes"] = articletypes


	this.Data["typename"] = typename
	this.Data["pageIndex"] = pageIndex
	this.Data["pageCount"] = int(pageCount)
	this.Data["count"] = count
	this.Data["articles"] = articles

	this.Layout = "layout.html"
	this.TplName = "index.html"
}






func (this *ArticleController)ShowAddArticle(){

	o:=orm.NewOrm()
	var articletypes []models.ArticleType
	o.QueryTable("ArticleType").All(&articletypes)
	this.Data["articletypes"] =articletypes
	Username := this.GetSession("username")
	this.Data["username"] = Username.(string)
	this.Layout = "layout.html"
	this.TplName = "add.html"
}
func (this*ArticleController)HandleAddArticle(){


	//获取数据
	aticleName := this.GetString("articleName")
	content := this.GetString("content")



//校验数据
if aticleName ==""||content == ""{
	this.Data["errmsg"] = "田间数据不完整"
	this.TplName = "add.html"
	return
}
file,head,err := this.GetFile("uploadname")
defer file.Close()


if err != nil{
	this.Data["errmsg"] = "文件上传失败"
	this.TplName = "add.html"
	return
}



if head.Size >5000000{
		this.Data["errmsg"] = "文件太大"
		this.TplName = "add.html"
		return
	}

	ext := path.Ext(head.Filename)

	if ext != ".jpg"&&ext!=".png"&&ext!=".jepg"{
		this.Data["errsmg"] = "文件格式错误"
		this.TplName = "add.html"
		return

	}
	fileName  := time.Now().Format("2006-01-02-15:04:05") + ext
	this.SaveToFile("uploadname","./static/img/"+fileName)


//数据库交互
	o:=orm.NewOrm()

	var article models.Article

	article.ArtiName = aticleName
	article.Acontent = content
	article.Aimg = "/static/img/"+fileName

	typename := this.GetString("select")

	var articletype  models.ArticleType
	articletype.Tname = typename

	o.Read(&articletype,"Tname")

	article.ArticleType = &articletype

	o.Insert(&article)


	//4.返回页面
	this.Redirect("/article/showArticlelist",302)
	}

func (this *ArticleController)ShowArticleDetail(){
	//抓取数据
	id,err := this.GetInt("articleId")
	if err != nil{
		beego.Info("连接传送错误")
	}
	o := orm.NewOrm()
	var article models.Article

	article.Id = id
	o.Read(&article)

	article.Acount += 1
	o.Update(&article)



	qs := o.QueryTable("Article")
	qs.RelatedSel("ArticleType").Filter("Id",id).All(&article)
	this.Data["article"] = article

	m2m := o.QueryM2M(&article,"Users")
	username := this.GetSession("username")
	if username == nil{
		this.Redirect("/loign",302)
		return
	}

	var user models.User
	user.Name = username.(string)
	o.Read(&user,"Name")
	m2m.Add(user)

var users []models.User
o.QueryTable("User").Filter("Articles__Article__Id",id).Distinct().All(&users)
this.Data["users"] = users

 Username:=this.GetSession("username")
 this.Data["username"] = Username.(string )
	this.Layout = "layout.html"
	this.TplName = "content.html"

}
func (this *ArticleController)ShowUpdateArticle(){
	id,err := this.GetInt("id")
	if err != nil{
		beego.Info("数据请求格式错误")
		this.Redirect("/article/showArticlelist",302)
		return
	}

	o := orm.NewOrm()

	var article models.Article

	article.Id = id

	err = o.Read(&article)
	if err != nil{
		beego.Info("id数据不存在")
		return
	}
	//article.ArtiName = Aticlename
	//article.Acontent  = Content
	//
	//o.Update(&article)
	Username := this.GetSession("username")
	this.Data["username"] = Username.(string)
	this.Data["article"]=article
	this.Layout = "layout.html"
	this.TplName = "update.html"


}

func UploadFile(filepath string,this beego.Controller,id int)string{
	file,head,err := this.GetFile(filepath)
	defer file.Close()
	if err != nil{
		beego.Info("图片上传错误")
		return ""
	}

	if head.Size > 50000000{
		beego.Info("文件上传过大")
		return ""
	}

	//o := orm.NewOrm()
	//var article models.Article
	//article.Id = id
	//o.Read(&article)



	ext := path.Ext(head.Filename)
	if head.Filename == ""{
		//return article.Aimg
	}else if ext != ".jpg" && ext != ".jpeg" && ext != ".npg"{

		beego.Info(head.Filename)
		beego.Info("文件上传格式不对")
		return ""
	}

	filename := time.Now().Format("2006-01-02-15:04:05")
	this.SaveToFile(filepath,"./static/img/"+filename+ext)

	return "/static/img/"+filename+ext
}
func (this *ArticleController)HandleUpdate(){

	id,err := this.GetInt("id")

	Content  := this.GetString("content")

	ArticleName := this.GetString("articleName")

	Filepath := UploadFile("uploadname",this.Controller,id )

	if err != nil|| ArticleName =="" ||Content == "" || Filepath == ""{
		beego.Info("数据格式错误")
		//this.Redirect("/UpdateAticle?id="+ strconv.Itoa(id),302)
		return
	}

	o:= orm.NewOrm()

	var article models.Article


	article.Id= id

	err = o.Read(&article)
	if err != nil{
		beego.Info("update id 不存在")
		//this.Redirect("/UpdateAticle?id="+ strconv.Itoa(id),302)
		return
	}


	article.ArtiName = ArticleName
	article.Acontent = Content
	article.Aimg = Filepath

	o.Update(&article)

	this.Redirect("/article/showArticlelist",302)

}

func (this *ArticleController)DeleteAticle(){
	id,err := this.GetInt("id")
	if err !=nil {
		beego.Info("id获取错误")
		return

	}


	o:=orm.NewOrm()

	var article models.Article

	article.Id = id

	o.Delete(&article)

	this.Redirect("/article/showArticlelist",302)

}

func(this *ArticleController)ShowArticleType(){
	o := orm.NewOrm()
	var articletypes []models.ArticleType

	o.QueryTable("ArticleType").All(&articletypes)


	this.Data["articletypes"] = articletypes
	Username := this.GetSession("username")
	this.Data["username"] = Username.(string)
	this.Layout = "layout.html"
	this.TplName= "addType.html"
}
func(this *ArticleController)HandleArticleType(){
	 typename  := this.GetString("typeName")
	 if typename == ""{
	 	beego.Info("信息不能为空")
	 	return
	 }
	 o := orm.NewOrm()
	 var articletype models.ArticleType
	 articletype.Tname = typename
	 o.Insert(&articletype)

	 this.Redirect("/article/addArticleType",302)
}
func(this *ArticleController)DeleteType(){
	id ,err:= this.GetInt("id")
	if err != nil{
		beego.Info("获取类型id失败")
		return
	}
	//获取orm对象
	o := orm.NewOrm()
	//获取处理数据对象
	var articletype models.ArticleType
	//绑定选取数据结构体值
	articletype.Id = id
	//删除选定结构体值
	o.Delete(&articletype)
//返回视图
this.Redirect("/article/addArticleType",302)
}