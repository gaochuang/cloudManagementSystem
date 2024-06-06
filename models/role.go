package models

type Role struct {
	Mode
	Name         string           `gorm:"column:name;comment:'role name';size:256" json:"name"`
	Desc         string           `gorm:"column:desc;comment:'role description';size:256" json:"desc"`
	Code         string           `gorm:"column:name;comment:'role code';size:256" json:"code"`
	Users        []User           `gorm:"foreignkey:RoleId" json:"users"`
	PermissionId int              `gorm:"column:permission_id;comment:'Permission id Foreign key'" json:"permission_id"`
	Permission   []RolePermission `gorm:"foreignkey:PermissionId" json:"permission"`
}

func (m Role) TableName() string {
	return m.Mode.TableName("role")
}
