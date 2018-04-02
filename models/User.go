package models

import (
	//"database/sql"

	"errors"
	//"fmt"
	//"github.com/astaxie/beedb"
	"github.com/astaxie/beego/orm"
	_ "github.com/ziutek/mymysql/godrv"
	//"fmt"
	"fmt"
)

type User struct {
	Id       int `PK`
	Username string
	Pwd      string
}

func (a *User) TableName() string {
	return UserTBName()
}

func SaveUser(user User) error {
	_,err := orm.NewOrm().Insert(&user)

	return err
}


func Validation(user User) error {
	o := orm.NewOrm()
	qs := o.QueryTable("rms_user")
	exist := qs.Filter("username",user.Username).Filter("pwd",user.Pwd).Exist()
	fmt.Printf("Is Exist: %s", exist)
	if exist {
		return nil
	}else {
		return errors.New("用户名或者密码错误")
	}

}


