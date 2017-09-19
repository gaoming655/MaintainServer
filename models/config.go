package models

import "github.com/astaxie/beego/logs"

type configservice struct {}

func (self *configservice)GetAll()([]*Config,error)  {
	var configs []*Config
	_,err := o.QueryTable("config").All(&configs)
	if err != nil{
		return nil,err
	}else{
		return configs,nil
	}
}

func (self *configservice)Save(c *Config)  {
	o.Insert(c)
}

func (self *configservice)GetbyId(cid int)*Config  {
	conf := new(Config)
	conf.Id = cid
	o.Read(conf)
	return conf

}

func (self *configservice)Update(c *Config) error  {
	logs.Debug(c)
	o.Update(c)
	return nil
}

func (self *configservice)DelById(id int)error  {
	conf := new(Config)
	conf.Id = id
	o.Delete(conf)
	return nil
}