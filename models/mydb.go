package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"net/url"
)



var (
	o orm.Ormer
	UserService *userservcie
	Auth  *myauth
	ConfigService  *configservice
)

func InitService()  {
	UserService = &userservcie{}
	Auth =  &myauth{}
	ConfigService = &configservice{}
}





func InitDB()  {
	dbHost := beego.AppConfig.String("db.host")
	dbPort := beego.AppConfig.String("db.port")
	dbUser := beego.AppConfig.String("db.user")
	dbPassword := beego.AppConfig.String("db.password")
	dbName := beego.AppConfig.String("db.name")
	timezone := beego.AppConfig.String("db.timezone")
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"
	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql",dsn,10,20)
	orm.RegisterModel(new(User),new(Config))
	orm.RunSyncdb("default",false,false)
	orm.Debug = true
	o = orm.NewOrm()
	orm.RunCommand()
}