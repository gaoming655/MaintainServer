package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"strconv"
	"MaintainServer/models"
	"MaintainServer/shortcut"
	"github.com/astaxie/beego/logs"
)

type BaseController struct {
	beego.Controller
}

func (self *BaseController)Prepare()  {
	self.Ctx.Output.Header("X-Author-By","GaoMing")
	self.InitAuth()

}

func (self *BaseController)InitAuth()  {
	token := self.Ctx.GetCookie("auth")
	controllerName, actionName := self.GetControllerAndAction()
	if !self.Verify(token) {
		if controllerName != "MainController" || (controllerName == "MainController" && actionName != "Login" && actionName != "Logout"){
			logs.Debug("跳转了?")
			self.Redirect(beego.URLFor("MainController.Login"),302)
		}
	}
}

func (self *BaseController)Verify(token string) bool {
	arr := strings.Split(token,"|")
	logs.Debug(len(arr))
	if len(arr) == 2{
		idstr,password := arr[0],arr[1]
		userid ,_ := strconv.Atoi(idstr)
		if userid >0 {
			user := models.UserService.GetUser(userid)
			if password == shortcut.Md5([]byte(user.Password)){
				logs.Debug("验证成功")
				return true
			} else {
				logs.Debug("Verify认证密码失败1")
				return false
			}
		}else{
			logs.Debug("Verify认证密码失败2")
			return false
		}
	}
	logs.Debug("Verify认证密码失败3")
	return false
}