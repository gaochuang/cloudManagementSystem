package user

import (
	"kubernetes_management_system/models"
)

//https://gorm.io/zh_CN/docs/models.html
//https://gorm.io/zh_CN/docs/conventions.html

type User struct {
	models.Mode
	UID      string `gorm:"column:uid;comment:'user uid'" json:"uid"`
	UserName string `gorm:"column:username;comment:'user name'" json:"userName"`
	Password string `gorm:"column:password;comment:'user password'" json:"password"`
	Status   *bool  `gorm:"type:tinyint(1);default:true;comment:'user status(enable/disable)'"`
}

type LoginUser struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (u *User) TableName() string {
	return u.Mode.TableName("users")
}
