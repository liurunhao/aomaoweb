package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"

	"xingaokao/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	username := c.GetSession("Username")  //获取session
	c.Data["Username"] = username

	c.TplName = "index.html"

}
func (c *MainController)ShowRegister(){
	c.TplName = "register.html"
}
func (c *MainController)ShowLogin(){
	c.TplName = "login.html"
}

//删除session
func (c *MainController) Logout() {
	c.DelSession("Username")
	response := models.ResponseJson{State:0,Message:"ok"}
	c.Data["json"] = response
	c.ServeJSON()
}
//登录验证和创建session
func (c *MainController) HandleLogin() {
	username:=c.GetString("username")
	password:=c.GetString("password")
	if username=="" || password=="" {
		beego.Info("用户名或密码不能为空")
		c.Redirect("/login",302)
		return
	}
	o := orm.NewOrm()
	user:=models.User{}
	user.Username=username
	user.Password=password
	if err:=o.Read(&user,"username");err!=nil {
		beego.Info("用户名输入错误，请重新输入")
		c.Redirect("/login",302)
		return
	}
	if user.Password!=password{
		beego.Info("密码错误，请重新输入")
		c.Redirect("/login",302)
		return
	}
	c.SetSession("Username",user.Username)
	c.Redirect("/",302)
}

func (c *MainController) HandleRegister() {
	username := c.GetString("username")
	password := c.GetString("password")
	if username == "" || password == "" {
		beego.Info("用户名或密码不能为空")
		c.Redirect("/register", 302)
		return
	}
	o := orm.NewOrm()
	user := models.User{}
	user.Username = username
	user.Password = password
	user.Created = time.Now().Unix()
	err := o.Read(&user, "username")
	if err != nil {
		user.Username = username
		user.Password = password
		_, e := o.Insert(&user)
		if e != nil {
			beego.Info("插入数据失败", e)
			return
		}
		beego.Info("注册成功！请登录！", err)
		c.Redirect("/login", 302)
		return
	}
	beego.Info("用户名已存在，请重新输入")
	c.Redirect("/register", 302)
}


func (c *MainController) Cart() {
	c.TplName = "cart.html"
}
func (c *MainController) CartAddress() {
	c.TplName = "cartAddress.html"
}
func (c *MainController) Orders() {
	c.TplName = "orders.html"
}



