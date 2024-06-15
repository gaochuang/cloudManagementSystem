package models

//https://gorm.io/zh_CN/docs/models.html
//https://gorm.io/zh_CN/docs/conventions.html

type User struct {
	Mode
	UID      string `gorm:"column:uid;comment:'user uid'" json:"uid"`
	UserName string `gorm:"column:username;comment:'user name'" json:"username",binding:"required"`
	Password string `gorm:"column:password;comment:'user password'" json:"password",binding:"required"`
	Status   *bool  `gorm:"type:tinyint(1);default:true;comment:'user status(enable/disable)'"`
	RoleId   uint   `gorm:"column:role_id;comment:'role foreign id'" json:"role_id"`
	Role     Role   `gorm:"foreignkey:RoleId" json:"role"`
}

type LoginUser struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type UsersChangePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type UsersListResponse struct {
	ID       int    `gorm:"column:id" json:"id"`
	UID      string `gorm:"column:uid" json:"uid"`
	UserName string `gorm:"column:username" json:"username"`
}

func (u *User) TableName() string {
	return u.Mode.TableName("users")
}
