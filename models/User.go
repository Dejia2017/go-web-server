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

//func getLink() beedb.Model {
//	//(*DB, error)
//	//连接数据的DSN格式为:
//	//username:password@protocol(address)/dbname?param=value
//	db, err := sql.Open("mysql", "root:518@tcp(127.0.0.1:3306)/GoLogin?charset=utf8")
//	if err != nil {
//		panic(err) //有错
//	}
//	orm := beedb.New(db)
//	return orm
//}

func SaveUser(user User) error {
	//fmt.Println(user.Username + user.Pwd + (string)(user.Id))
	//o := orm.NewOrm()
	//o.Insert()
	_,err := orm.NewOrm().Insert(&user)
	//fmt.Println("Insert user' data success! \r\n")
	//if num, err := orm.NewOrm().Insert(user); err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("Insert %d user' data!\r\n", num)
	//}
	//fmt.Println(o.Insert)
	//o.Insert()
	//m := User{}

	//fmt.Println(db)

	//err := orm.Save(&user)
	//err := beedb.New(orm.GetDB()).Save(&user)
	//orm := getLink()
	//fmt.Println(user)
	//err := orm.Save(&user)
	return err
}

//func (c *BackendUserController) Save() {
//	m := models.BackendUser{}
//	o := orm.NewOrm()
//	var err error
//	//获取form里的值
//	if err = c.ParseForm(&m); err != nil {
//		c.jsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
//	}
//	//删除已关联的历史数据
//	if _, err := o.QueryTable(models.RoleBackendUserRelTBName()).Filter("backenduser__id", m.Id).Delete(); err != nil {
//		c.jsonResult(enums.JRCodeFailed, "删除历史关系失败", "")
//	}
//	if m.Id == 0 {
//		//对密码进行加密
//		m.UserPwd = utils.String2md5(m.UserPwd)
//		if _, err := o.Insert(&m); err != nil {
//			c.jsonResult(enums.JRCodeFailed, "添加失败", m.Id)
//		}
//	} else {
//		if oM, err := models.BackendUserOne(m.Id); err != nil {
//			c.jsonResult(enums.JRCodeFailed, "数据无效，请刷新后重试", m.Id)
//		} else {
//			m.UserPwd = strings.TrimSpace(m.UserPwd)
//			if len(m.UserPwd) == 0 {
//				//如果密码为空则不修改
//				m.UserPwd = oM.UserPwd
//			} else {
//				m.UserPwd = utils.String2md5(m.UserPwd)
//			}
//			//本页面不修改头像和密码，直接将值附给新m
//			m.Avatar = oM.Avatar
//		}
//		if _, err := o.Update(&m); err != nil {
//			c.jsonResult(enums.JRCodeFailed, "编辑失败", m.Id)
//		}
//	}
//	//添加关系
//	var relations []models.RoleBackendUserRel
//	for _, roleId := range m.RoleIds {
//		r := models.Role{Id: roleId}
//		relation := models.RoleBackendUserRel{BackendUser: &m, Role: &r}
//		relations = append(relations, relation)
//	}
//	if len(relations) > 0 {
//		//批量添加
//		if _, err := o.InsertMulti(len(relations), relations); err == nil {
//			c.jsonResult(enums.JRCodeSucc, "保存成功", m.Id)
//		} else {
//			c.jsonResult(enums.JRCodeFailed, "保存失败", m.Id)
//		}
//	} else {
//		c.jsonResult(enums.JRCodeSucc, "保存成功", m.Id)
//	}
//}

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


