package user

import "gorm.io/gorm"

//https://gorm.io/zh_CN/docs/models.html
//https://gorm.io/zh_CN/docs/conventions.html

type User struct {
	gorm.Model
	UID      string `gorm:"column:uid;comment:'user id'" json:"uid"`
	UserName string `gorm:"column:username;comment:'user name'" json:"userName"`
	Password string `gorm:"column:password;comment:'user password'" json:"password"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) TableName() string {
	return "users"
}
