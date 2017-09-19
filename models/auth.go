package models

import (
	"MaintainServer/shortcut"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/pkg/errors"
	"strconv"
)

type myauth struct {}

func (self *myauth)Login(username,password string) (string,error) {
	err,user := UserService.GetUserByName(username)
	if err != nil{
		logs.Debug("用户不存在")
		return "",errors.New("用户不在")
	}else{
		if user.Password != shortcut.Md5([]byte(password)){
			logs.Debug("密码不正确")
			return "",errors.New("密码不正确")
		}else {
			idstr := strconv.Itoa(user.Id)
			token := fmt.Sprintf("%s|%s",idstr,shortcut.Md5([]byte(user.Password)))
			return token,nil
		}
	}
}