package models

type User struct {
	Id int
	Username string `orm:"size(10)"`
	Password string `orm:"size(40)"`
}

type Config struct {
	Id int      `json:"-" form:"-"`
	Host string `orm:"size(50)" json:"host" form:"host" valid:"Required"`   //域
	Result string `orm:"type(text)" json:"result" form:"result" valid:"Required"`  //返回结果
	Ps string `orm:"size(30)" json:"-" form:"ps" valid:"Required;MaxSize(30)"`   //备注
	Method string `orm:"size(10)" json:"method" form:"method" valid:"Required"` //请求方法
}

