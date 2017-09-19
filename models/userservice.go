package models

type userservcie struct {}

func (self *userservcie)GetUser(id int) *User {
	user := &User{}
	user.Id = id
	o.Read(user)
	return user
}

func (self *userservcie)GetUserByName(name string) (error ,*User)  {
	user := &User{}
	user.Username = name
	err := o.Read(user,"Username")
	if err != nil{
		return err,nil
	}
	return nil,user
}