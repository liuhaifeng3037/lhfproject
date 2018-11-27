package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type TestController struct {
	beego.Controller
}

type Test2Controller struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
func (c *TestController) Get() {
	c.TplName = "testindex.html"
}

func (c *Test2Controller) Get() {
	c.TplName = "index1.tpl"
}

