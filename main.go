package main

import (

	_ "dddd/routers"
	"github.com/astaxie/beego"
	_"dddd/models"


)

func main() {
	beego.AddFuncMap("prepage",ShowPrePage)
	beego.AddFuncMap("nextpage",ShowNextPage)
	beego.Run()
}
func ShowPrePage(pageIndex int)int{
	if pageIndex == 1{
		return pageIndex
	}
	return pageIndex -1
}
func ShowNextPage(pageIndex int,pageCount int)int{
	if pageIndex == pageCount{
		return pageIndex
	}
	return pageIndex +1
}