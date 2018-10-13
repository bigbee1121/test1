package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"time"
)
type User struct {
	Id int
	Name string `orm:"unique"`
	Password string `orm:"size(20)"`
	Articles []*Article `orm:"rel(m2m)"` //设置多对多关系
}
//文章结构体
type Article struct {
	Id int `orm:"pk;auto"`
	ArtiName string `orm:"size(20)"`
	Atime time.Time `orm:"auto_now"`
	Acount int `orm:"default(0);null"`
	Acontent string `orm:"size(500)"`
	Aimg string `orm:"size(100)"`

	ArticleType*ArticleType `orm:"rel(fk);null;on_delete(set_null)"` //设置一对多关系
	Users []*User `orm:"reverse(many)"` //设置多对多的反向关系
}
//类型表
type ArticleType struct {
	Id int
	Tname string `orm:"size(20)"`
	Articles []*Article `orm:"reverse(many)"` //设置一对多的反向关系
}
func init(){
	orm.RegisterDataBase("default","mysql","root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	orm.RegisterModel(new(User),new(Article),new(ArticleType))
	orm.RunSyncdb("default",false,true )
}
