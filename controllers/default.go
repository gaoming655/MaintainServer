package controllers

import (
	"github.com/astaxie/beego"
	"MaintainServer/models"
	"github.com/astaxie/beego/validation"
	"github.com/astaxie/beego/logs"
	"strconv"
)

type MainController struct {
	BaseController
}


func (c *MainController)Logout()  {
	c.Ctx.SetCookie("auth","")
	c.Redirect(beego.URLFor("MainController.Login"),302)
}


func (c *MainController) Login() {
	beego.ReadFromRequest(&c.Controller)
	if c.Ctx.Request.Method == "POST"{
		flash := beego.NewFlash()
		username := c.GetString("username")
		password := c.GetString("password")
		if username != "" && password != "" {
			token, err := models.Auth.Login(username, password)
			if err != nil {
				flash.Error(err.Error())
				flash.Store(&c.Controller)
				c.Redirect(beego.URLFor("MainController.Login"),302)
			} else {
				c.Ctx.SetCookie("auth",token,3600)
				c.Redirect(beego.URLFor("MainController.List"),302)
			}
		}


	}
	c.TplName = "login.tpl"
}

func (c *MainController)List()  {
	table,_ := models.ConfigService.GetAll()
	c.Data["table"] = table
	c.TplName="report.html"
}

func (c *MainController)Add()  {
	flush := beego.ReadFromRequest(&c.Controller)
	c.Data["maps"] = flush.Data
	c.Data["title"] = "添加规则"
	if c.Ctx.Request.Method == "POST" {
		flush := beego.NewFlash()
		config := new(models.Config)
		valid := validation.Validation{}
		if err := c.ParseForm(config); err != nil {
			c.Ctx.WriteString("error1")

		} else {
			b, _ := valid.Valid(config)
			if !b {
				for _, err := range valid.Errors {
					flush.Set(err.Key,err.Message)
					logs.Debug(err.Key, err.Message)
				}
				flush.Set("Host",config.Host)
				flush.Set("Ps",config.Ps)
				flush.Set("Result",config.Result)
				flush.Store(&c.Controller)
				c.Redirect(beego.URLFor("MainController.Add"),302)
			}else{
				models.ConfigService.Save(config)
				c.Redirect(beego.URLFor("MainController.List"),302)
			}
		}
	}
	c.TplName = "add.html"
}

func (c *MainController)Edit()  {
	flush := beego.ReadFromRequest(&c.Controller)
	qid := c.Ctx.Input.Param(":id")
	cid,_ := strconv.Atoi(qid)
	c.Data["maps"] = flush.Data
	c.Data["title"] = "修改规则"

	if c.Ctx.Request.Method == "POST"{
		flush := beego.NewFlash()
		config := new(models.Config)
		valid := validation.Validation{}
		if err := c.ParseForm(config); err != nil {
			c.Ctx.WriteString("error1")

		} else {
			b, _ := valid.Valid(config)
			if !b {
				for _, err := range valid.Errors {
					flush.Set(err.Key,err.Message)
					logs.Debug(err.Key, err.Message)
				}
				flush.Set("Host",config.Host)
				flush.Set("Ps",config.Ps)
				flush.Set("Result",config.Result)
				flush.Set("Method",config.Method)
				flush.Store(&c.Controller)
				c.Redirect(beego.URLFor("MainController.Edit",":id",cid),302)
			}else{
				config.Id =cid
				//UPdate config 需要修改
				models.ConfigService.Update(config)
				c.Redirect(beego.URLFor("MainController.List"),302)
			}
		}
	}else {
		config := models.ConfigService.GetbyId(cid)
		c.Data["form"] = map[string]string{"Host":config.Host,"Ps":config.Ps,"Result":config.Result,"Method":config.Method}
		c.TplName = "edit.html"
	}

}

func (c *MainController)Del()  {
	qid := c.Ctx.Input.Param(":id")
	cid ,_ := strconv.Atoi(qid)
	models.ConfigService.DelById(cid)
	c.Redirect(beego.URLFor("MainController.List"),302)


}




// api 无需验证

type ApiController struct {
	beego.Controller
}


func (c *ApiController)Api()  {
	table,_ := models.ConfigService.GetAll()
	c.Data["json"] = table
	c.ServeJSON()
}

